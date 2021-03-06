import React, { useEffect, useMemo, useRef, useState } from "react";
import { Grid, TextField, Typography } from "@material-ui/core";
import LocationOnIcon from "@material-ui/icons/LocationOn";
import Autocomplete from "@material-ui/lab/Autocomplete";
import parse from "autosuggest-highlight/parse";
import throttle from "lodash/throttle";
import PropTypes from "prop-types";

const googleMapsKey = process.env.REACT_APP_GOOGLE_MAPS_API_KEY; // fill this in with your own key

function loadScript(src, position, id) {
  if (!position) {
    return;
  }

  const script = document.createElement("script");
  script.setAttribute("async", "");
  script.setAttribute("id", id);
  script.src = src;
  position.appendChild(script);
}

const geocoderService = { current: null };
const autocompleteService = { current: null };
const placesService = { current: null };

if (!googleMapsKey && !process.env.IN_TEST) {
  // eslint-disable-next-line no-console
  console.warn(`No Google API key specified.
Run \`REACT_APP_GOOGLE_MAPS_API_KEY=<api key> yarn start\` to have location-based features work.`);
}

export function useGoogleMaps(dependencies) {
  const loaded = useRef(false);

  if (typeof window !== "undefined" && googleMapsKey && !loaded.current) {
    if (!document.querySelector("#google-maps")) {
      loadScript(
        `https://maps.googleapis.com/maps/api/js?key=${googleMapsKey}&libraries=places`,
        document.querySelector("head"),
        "google-maps"
      );
    }

    loaded.current = true;
  }

  useEffect(() => {
    if (typeof window !== "undefined" && window.google) {
      if (!geocoderService.current) {
        geocoderService.current = new window.google.maps.Geocoder();
      }
      if (!autocompleteService.current) {
        autocompleteService.current = new window.google.maps.places.AutocompleteService();
      }
      if (!placesService.current) {
        const el = document.createElement("div");
        placesService.current = new window.google.maps.places.PlacesService(el);
      }
    }
  }, dependencies);

  return useMemo(
    () => ({
      isReady: () =>
        Boolean(
          geocoderService.current &&
            autocompleteService.current &&
            placesService.current
        ),
      geocode: (req, cb) => geocoderService.current.geocode(req, cb),
      getPlacePredictions: throttle(
        (req, cb) => autocompleteService.current.getPlacePredictions(req, cb),
        200
      ),
      getPlaceDetails: (req, cb) => placesService.current.getDetails(req, cb),
    }),
    []
  );
}

/**
 * A text field where users type the name of a physical location and select from a list of
 * matching locations found on Google Maps.
 *
 * Used to pick meetup location in [CreateMeetup](#createmeetup) and [Search](#search).
 */
export default function LocationPicker({
  value,
  setValue,
  style,
  disabled = false,
}) {
  const [rawValue, setRawValue] = useState(null); // AutocompletePrediction
  const [inputValue, setInputValue] = useState("");
  const [options, setOptions] = useState([]);

  const { getPlacePredictions, getPlaceDetails } = useGoogleMaps([
    value,
    inputValue,
  ]);

  useEffect(() => {
    const controller = new AbortController();

    if (rawValue) {
      getPlaceDetails(
        {
          placeId: rawValue.place_id,
          fields: ["geometry.location"],
        },
        (result) => {
          if (!controller.signal.aborted && result) {
            setValue({
              description: rawValue.description,
              coords: [
                result.geometry.location.lat(),
                result.geometry.location.lng(),
              ],
            });
          }
        }
      );
    }

    return () => {
      controller.abort();
    };
  }, [rawValue]);

  useEffect(() => {
    if (!autocompleteService.current) {
      return undefined; // Google Maps not initialized yet
    }

    if (inputValue === "") {
      setOptions(value ? [value] : []);
      return undefined;
    }

    const controller = new AbortController();

    getPlacePredictions({ input: inputValue }, (results) => {
      if (!controller.signal.aborted) {
        const newOptions = [];
        if (value) {
          newOptions.push(value);
        }
        if (results) {
          newOptions.push(...results);
        }
        setOptions(newOptions);
      }
    });

    return () => {
      controller.abort();
      if (getPlacePredictions.cancel) {
        getPlacePredictions.cancel();
      }
    };
  }, [value, inputValue, getPlacePredictions]);

  return (
    <Autocomplete
      disabled={disabled}
      style={style}
      getOptionLabel={(option) =>
        typeof option === "string" ? option : option.description
      }
      filterOptions={(x) => x}
      options={options}
      autoComplete
      includeInputInList
      filterSelectedOptions
      value={value}
      onChange={(event, newValue) => {
        setOptions(newValue ? [newValue, ...options] : options);
        setRawValue(newValue);
      }}
      onInputChange={(event, newInputValue) => {
        setInputValue(newInputValue);
      }}
      renderInput={(params) => (
        <TextField
          {...params} // eslint-disable-line react/jsx-props-no-spreading
          required
          label="Add a location"
          variant="outlined"
          fullWidth
        />
      )}
      renderOption={(option) => {
        const matches =
          option.structured_formatting.main_text_matched_substrings;
        const parts = parse(
          option.structured_formatting.main_text,
          matches.map((match) => [match.offset, match.offset + match.length])
        );

        return (
          <Grid container alignItems="center">
            <Grid item>
              <LocationOnIcon style={{ margin: 8 }} />
            </Grid>
            <Grid item xs>
              {parts.map((part) => (
                <span
                  key={part.text}
                  style={{ fontWeight: part.highlight ? 700 : 400 }}
                >
                  {part.text}
                </span>
              ))}

              <Typography variant="body2" color="textSecondary">
                {option.structured_formatting.secondary_text}
              </Typography>
            </Grid>
          </Grid>
        );
      }}
    />
  );
}

LocationPicker.propTypes = {
  /** Location value for physical locations. */
  value: PropTypes.shape({
    /** Name of the location, e.g. `"UCLA"`. May be empty for unnamed locations. */
    description: PropTypes.string.isRequired,
    /** An ordered pair `[latitude, longitude]`, e.g. `[ 34.069107615481094, -118.44521328860678]`. */
    coords: PropTypes.arrayOf(PropTypes.number.isRequired).isRequired,
  }),
  /** Setter for the location value, generated from useEffect(). */
  setValue: PropTypes.func.isRequired,
  /** Styling for the Autocomplete component (see [Material UI - Autocomplete](https://material-ui.com/components/autocomplete/)).
   */
  style: PropTypes.objectOf(
    PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired
  ),
  disabled: PropTypes.bool,
};

LocationPicker.defaultProps = {
  value: null,
  style: null,
  disabled: false,
};
