import React, { useEffect, useState } from "react";
import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  Select,
  Typography,
} from "@material-ui/core";
import { Link } from "react-router-dom";

import LocationPicker, { useGoogleMaps } from "../common/LocationPicker";
import MeetupCard from "../common/MeetupCard";
import { IN_PERSON, REMOTE } from "../../constants";
import * as fetcher from "../../lib/fetch";
import TagPicker from "../common/TagPicker";

/**
 * Search/home page. Users can choose to search for in-person or remote/online meetups
 * according to their interests. If in-person meetups are chosen, a
 * [LocationPicker](#locationpicker) is shown.
 *
 * Results are shown as a list of [MeetupCards](#meetupcard). Only meetups that
 * have not passed and have not been cancelled are returned.
 */
function Search() {
  const [error, setError] = useState(null);
  const [tags, setTags] = useState([]);
  const [coords, setCoords] = useState(null);
  const [position, setPosition] = useState(null); // position object from navigator.geolocation
  const [meetups, setMeetups] = useState(null);
  const [meetupType, setMeetupType] = useState("");
  const [searchLocation, setSearchLocation] = useState(null);
  const [radius, setRadius] = useState(1);

  const resetSearch = () => {
    setError(null);
    setTags([]);
    setCoords(null);
    setPosition(null);
    setMeetups(null);
    setMeetupType("");
    setSearchLocation(null);
    setRadius("");
  };

  const validateSearch = () => {
    if (meetupType === "") {
      return false;
    }
    if (
      meetupType === IN_PERSON &&
      (searchLocation === null || radius === "")
    ) {
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
        lat: searchLocation.coords[0],
        lon: searchLocation.coords[1],
        radius,
        tags,
      });
      if (res.ok) {
        setMeetups(resJSON.filter((meetup) => !meetup.canceled));
      }
    } else {
      const { res, resJSON } = await fetcher.searchForRemoteMeetups(tags);
      if (res.ok) {
        setMeetups(resJSON.filter((meetup) => !meetup.canceled));
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
        <FormControl
          required
          variant="outlined"
          style={{ width: 120, marginRight: 20 }}
        >
          <InputLabel htmlFor="select-meetup-type">Type</InputLabel>
          <Select
            label="Type"
            native
            id="select-meetup-type"
            value={meetupType}
            onChange={(event) => setMeetupType(event.target.value)}
          >
            <option aria-label="None" value="" />
            <option value={IN_PERSON}>In person</option>
            <option value={REMOTE}>Remote</option>
          </Select>
        </FormControl>
        <TagPicker tags={tags} setTags={setTags} />
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
          <InputLabel htmlFor="select-radius">Radius</InputLabel>
          <Select
            label="Radius"
            native
            id="select-radius"
            value={radius}
            onChange={(event) => setRadius(event.target.value)}
          >
            <option value={1}>1 km</option>
            <option value={5}>5 km</option>
            <option value={10}>10 km</option>
            <option value={50}>50 km</option>
            <option value={100}>100 km</option>
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
        owner={meetup.owner}
        tags={meetup.tags}
        canceled={meetup.canceled}
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
        <Box alignSelf="flex-end" mt={2} mb={2}>
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
