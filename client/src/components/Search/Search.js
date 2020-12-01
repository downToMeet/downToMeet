import React, { useEffect, useState } from "react";
import AutoComplete from "@material-ui/lab/Autocomplete";
import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  TextField,
  Typography,
} from "@material-ui/core";
import { Link } from "react-router-dom";

import LocationPicker, { useGoogleMaps } from "../common/LocationPicker";
import MeetupCard from "./MeetupCard";
import { IN_PERSON, REMOTE } from "../../constants";
import * as fetcher from "../../lib/fetch";

function Search() {
  const [error, setError] = useState(null);
  const [tags, setTags] = useState([]);
  const [coords, setCoords] = useState(null);
  const [position, setPosition] = useState(null); // position object from navigator.geolocation
  const [meetups, setMeetups] = useState(null);
  const [meetupType, setMeetupType] = useState("");
  const [searchLocation, setSearchLocation] = useState(null);
  const [radius, setRadius] = useState(null);

  // TODO: load options from the server instead
  const tagOptions = [
    "basketball",
    "ping pong",
    "badminton",
    "movie",
    "cooking",
  ];

  const resetSearch = () => {
    setError(null);
    setTags([]);
    setCoords(null);
    setPosition(null);
    setMeetups(null);
    setMeetupType("");
    setSearchLocation(null);
  };

  const validateSearch = () => {
    if (meetupType === "") {
      return false;
    }
    if ((meetupType === IN_PERSON && coords === null) || radius === null) {
      return false;
    }
    return true;
  };

  const onSubmit = async () => {
    if (!validateSearch()) {
      setError(true);
      return;
    }
    setError(false);
    if (meetupType === IN_PERSON) {
      const { res, resJSON } = await fetcher.searchForMeetups({
        lat: coords[0],
        lon: coords[1],
        radius,
        tags,
      });
      if (res.ok) {
        setMeetups(resJSON);
      }
    } else {
      const { res, resJSON } = await fetcher.searchForRemoteMeetups(tags);
      if (res.ok) {
        setMeetups(resJSON);
      }
    }
  };

  const locate = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((pos) => {
        const { latitude, longitude } = pos.coords;
        setCoords([latitude, longitude]);
        setPosition(pos);
      });
    }
  };

  const { isReady, geocode } = useGoogleMaps([position]);

  useEffect(() => {
    // Try to geocode the current user location – but only if the current
    // coordinates to search for are the ones we got from geolocation.
    if (
      position &&
      position.coords.latitude === coords[0] &&
      position.coords.longitude === coords[1] &&
      isReady()
    ) {
      const controller = new AbortController();

      geocode(
        {
          location: {
            lat: position.coords.latitude,
            lng: position.coords.longitude,
          },
        },
        (res) => {
          // If this request has been aborted, or if the current coordinates to
          // search for are no longer the position we got from geolocation,
          // abort.
          if (
            controller.signal.aborted ||
            position.coords.latitude !== coords[0] ||
            position.coords.longitude !== coords[1]
          ) {
            return;
          }

          for (const place of res || []) {
            // Don't try to get accurate places if the location we got had low
            // accuracy.
            if (
              position.coords.accuracy > 2000 &&
              place.geometry.location_type !== "APPROXIMATE"
            ) {
              continue;
            }

            setSearchLocation({
              description: place.formatted_address,
              coords: [position.coords.latitude, position.coords.longitude],
            });
            break;
          }
        }
      );
      return () => {
        controller.abort();
      };
    }
    return undefined;
  }, [position]);

  const renderSearch = () => {
    return (
      <Box display="flex" flexDirection="row" flexWrap="wrap" mt={2}>
        <FormControl required variant="outlined" style={{ width: 120 }}>
          <InputLabel id="select-meetup-type-label">Type</InputLabel>
          <Select
            label="Type"
            labelId="select-meetup-type-label"
            value={meetupType}
            onChange={(event) => setMeetupType(event.target.value)}
          >
            <MenuItem value={IN_PERSON}>In person</MenuItem>
            <MenuItem value={REMOTE}>Remote</MenuItem>
          </Select>
        </FormControl>
        <AutoComplete
          multiple
          disableCloseOnSelect
          value={tags}
          onChange={(event, newValue) => setTags(newValue)}
          variant="outlined"
          options={tagOptions}
          renderInput={(params) => (
            <TextField
              // eslint-disable-next-line react/jsx-props-no-spreading
              {...params}
              variant="outlined"
              label="Search for your interests"
            />
          )}
          style={{ marginLeft: 20, flex: 1 }}
        />
      </Box>
    );
  };

  const renderLocation = () => {
    return (
      <Box display="flex" alignItems="center" mt={2}>
        {navigator.geolocation && (
          <>
            <Button onClick={locate} variant="contained" style={{ width: 120 }}>
              Use GPS
            </Button>
            <Typography style={{ margin: 10 }}>or</Typography>
          </>
        )}
        <LocationPicker
          value={searchLocation}
          setValue={setSearchLocation}
          style={{
            flex: 1,
          }}
        />
        <FormControl
          required
          variant="outlined"
          style={{ width: 100, marginLeft: 10 }}
        >
          <InputLabel id="select-radius-label">Radius</InputLabel>
          <Select
            label="Radius"
            labelId="select-radius-label"
            value={radius}
            onChange={(event) => setRadius(event.target.value)}
          >
            <MenuItem value={1}>1 km</MenuItem>
            <MenuItem value={5}>5 km</MenuItem>
            <MenuItem value={10}>10 km</MenuItem>
            <MenuItem value={50}>50 km</MenuItem>
            <MenuItem value={100}>100 km</MenuItem>
          </Select>
        </FormControl>
      </Box>
    );
  };

  const renderMeetups = () => {
    if (meetups === null) {
      return null;
    }
    if (meetups.length === 0) {
      return (
        <Typography style={{ marginTop: 16 }}>
          We couldn’t find any meetups, try widening your search or{" "}
          <Link to="/create">create your own meetup!</Link>
        </Typography>
      );
    }
    const meetupCards = meetups.map((meetup) => (
      <MeetupCard
        key={meetup.id}
        title={meetup.title}
        time={meetup.time}
        location={meetup.location}
        id={meetup.id}
        organizer={meetup.organizer}
        tags={meetup.tags}
      />
    ));
    return meetupCards;
  };

  return (
    <Container maxWidth="sm">
      <Typography
        variant="h2"
        component="h1"
        style={{ marginTop: 20, textAlign: "center" }}
      >
        Search for Meetups
      </Typography>
      {error && (
        <Typography color="error">
          Please ensure all required fields (marked with *) are filled out.
        </Typography>
      )}
      <Box display="flex" flexDirection="column">
        {renderSearch()}
        {meetupType === IN_PERSON && renderLocation()}
        <Box alignSelf="flex-end" mt={2}>
          <Button
            onClick={resetSearch}
            variant="contained"
            style={{ marginRight: 20 }}
          >
            Reset Search
          </Button>
          <Button onClick={onSubmit} variant="contained">
            Search
          </Button>
        </Box>
        {renderMeetups()}
      </Box>
    </Container>
  );
}

export default Search;
