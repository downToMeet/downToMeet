import React, { useState } from "react";
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

import LocationPicker from "../common/LocationPicker";
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
  const [lat, setLat] = useState(null);
  const [lon, setLon] = useState(null);
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
      lat,
      lon,
      radius: 10,
      tags,
    });
    if (res.ok) {
      setMeetups(resJSON);
    }
  };

  const locate = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition((position) => {
        const { latitude, longitude } = position.coords;
        setLat(latitude);
        setLon(longitude);
      });
    }
  };

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
        <Button onClick={locate} variant="contained">
          Use GPS
        </Button>
        <Typography style={{ margin: 10 }}>or</Typography>
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
