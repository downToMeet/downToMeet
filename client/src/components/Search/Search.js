import React, { useEffect, useState } from "react";
import AutoComplete from "@material-ui/lab/Autocomplete";
import {
  Box,
  Button,
  Container,
  TextField,
  Typography,
} from "@material-ui/core";
import makeStyles from "@material-ui/styles/makeStyles";
import { Link } from "react-router-dom";

import LocationPicker, { useGoogleMaps } from "../common/LocationPicker";
import MeetupCard from "./MeetupCard";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles(() => ({
  searchBar: {
    marginTop: 40,
  },
}));

function Search() {
  const classes = useStyles();

  const [tags, setTags] = useState([]);
  const [coords, setCoords] = useState(null);
  const [position, setPosition] = useState(null); // position object from navigator.geolocation
  const [meetups, setMeetups] = useState(null);
  const [searchLocation, setSearchLocation] = useState(null);

  // TODO: load options from the server instead
  const tagOptions = [
    "basketball",
    "ping pong",
    "badminton",
    "movie",
    "cooking",
  ];

  const onSubmit = async () => {
    const { res, resJSON } = await fetcher.searchForMeetups({
      lat: coords[0],
      lon: coords[1],
      radius: 10,
      tags,
    });
    if (res.ok) {
      setMeetups(resJSON);
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

  const renderSearchBar = () => {
    return (
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
            className={classes.searchBar}
          />
        )}
      />
    );
  };

  const renderLocation = () => {
    return (
      <Box display="flex" alignItems="center">
        {navigator.geolocation && (
          <>
            <Button onClick={locate} variant="contained">
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
      </Box>
    );
  };

  const renderMeetups = () => {
    if (meetups === null) {
      return null;
    }
    if (meetups.length === 0) {
      return (
        <Typography>
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
      {renderSearchBar()}
      {renderLocation()}
      <Button onClick={onSubmit} variant="contained">
        Search
      </Button>
      {renderMeetups()}
    </Container>
  );
}

export default Search;
