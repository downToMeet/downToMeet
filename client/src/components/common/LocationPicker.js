import React, { useEffect, useMemo, useRef, useState } from "react";
import { Grid, TextField, Typography } from "@material-ui/core";
import LocationOnIcon from "@material-ui/icons/LocationOn";
import Autocomplete from "@material-ui/lab/Autocomplete";
import parse from "autosuggest-highlight/parse";
import throttle from "lodash/throttle";
import PropTypes from "prop-types";

const googleMapsKey = "AIzaSyCUhoOofhPE1Cc7jVRE2-Lfri_ueCiDcDM"; // fill this in with your own key

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

const autocompleteService = { current: null };
const placesService = { current: null };

function useGoogleMaps(dependencies) {
  const loaded = useRef(false);

  if (typeof window !== "undefined" && !loaded.current) {
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
      getPlacePredictions: throttle(
        (req, cb) => autocompleteService.current.getPlacePredictions(req, cb),
        200
      ),
      getPlaceDetails: (req, cb) => placesService.current.getDetails(req, cb),
    }),
    []
  );
}

export default function LocationPicker({ value, setValue, style }) {
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
  value: PropTypes.shape({
    description: PropTypes.string.isRequired,
    coords: PropTypes.arrayOf(PropTypes.number.isRequired).isRequired,
  }),
  setValue: PropTypes.func.isRequired,
  style: PropTypes.objectOf(
    PropTypes.oneOfType([PropTypes.string, PropTypes.number]).isRequired
  ),
};

LocationPicker.defaultProps = {
  value: null,
  style: null,
};
