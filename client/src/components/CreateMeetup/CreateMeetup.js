import React, { useState } from "react";
import makeStyles from "@material-ui/styles/makeStyles";
import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  Select,
  Slider,
  TextField,
  Typography,
} from "@material-ui/core";
import AutoComplete from "@material-ui/lab/Autocomplete";
import {
  KeyboardDateTimePicker,
  MuiPickersUtilsProvider,
} from "@material-ui/pickers";
import DayUtils from "@date-io/dayjs";

import LocationPicker from "../common/LocationPicker";
import { IN_PERSON, REMOTE } from "../../constants";
import * as fetcher from "../../lib/fetch";

const useStyles = makeStyles(() => ({
  formSection: {
    marginTop: 15,
    marginBottom: 15,
  },
  formInput: {
    width: 200,
  },
}));

function CreateMeetup() {
  const classes = useStyles();
  const [title, setTitle] = useState("");
  const [time, setTime] = useState(new Date());
  const [meetupType, setMeetupType] = useState("");
  const [meetupURL, setMeetupURL] = useState("");
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([2, 10]);
  const [description, setDescription] = useState("");
  const [tags, setTags] = useState([]);

  const [error, setError] = useState(false);
  const [creatingMeetup, setCreatingMeetup] = useState(false);

  const clearForm = () => {
    setTitle("");
    setTime(new Date());
    setMeetupType("");
    setMeetupURL("");
    setMeetupLocation(null);
    setGroupCount([2, 10]);
    setDescription("");
    setTags([]);
    setError(false);
    setCreatingMeetup(false);
  };

  // TODO: load options from the server instead
  const tagOptions = [
    "basketball",
    "ping pong",
    "badminton",
    "movie",
    "cooking",
  ];

  const validateForm = () => {
    if (title === "" || meetupType === "" || tags.length === 0) {
      setError(true);
      return false;
    }
    if (meetupType === IN_PERSON && meetupLocation === null) {
      setError(true);
      return false;
    }
    if (meetupType === REMOTE && meetupURL === "") {
      setError(true);
      return false;
    }
    setError(false);
    return true;
  };

  const renderNameInput = () => {
    return (
      <div className={classes.formSection}>
        <TextField
          required
          variant="outlined"
          label="Title"
          value={title}
          onChange={(event) => setTitle(event.target.value)}
          style={{ width: "100%" }}
        />
      </div>
    );
  };

  const renderDateInput = () => {
    return (
      <div className={classes.formSection}>
        <MuiPickersUtilsProvider utils={DayUtils}>
          <KeyboardDateTimePicker
            required
            variant="inline"
            inputVariant="outlined"
            label="Time"
            value={time}
            onChange={(newDate) => setTime(newDate)}
            className={classes.formInput}
            style={{
              width: "100%",
            }}
          />
        </MuiPickersUtilsProvider>
      </div>
    );
  };

  const renderMeetupLocation = () => {
    return (
      <Box display="flex" flexWrap="wrap" className={classes.formSection}>
        <FormControl required variant="outlined" style={{ width: 150 }}>
          <InputLabel id="select-meetup-type-label">Type</InputLabel>
          <Select
            label="Type"
            data-testid="select-meetup-type"
            labelId="select-meetup-type-label"
            value={meetupType}
            native
            onChange={(event) => setMeetupType(event.target.value)}
          >
            <option aria-label="None" value="" />
            <option value={IN_PERSON}>In person</option>
            <option value={REMOTE}>Remote</option>
          </Select>
        </FormControl>
        {meetupType === IN_PERSON && (
          <LocationPicker
            value={meetupLocation}
            setValue={setMeetupLocation}
            style={{
              marginLeft: 20,
              flex: 1,
            }}
          />
        )}
        {meetupType === REMOTE && (
          <TextField
            label="URL"
            variant="outlined"
            required
            value={meetupURL}
            onChange={(event) => setMeetupURL(event.target.value)}
            className={classes.formInput}
            style={{
              marginLeft: 20,
              flex: 1,
            }}
          />
        )}
      </Box>
    );
  };

  const renderGroupSizeInput = () => {
    return (
      <Box
        display="flex"
        justifyContent="space-between"
        className={classes.formSection}
      >
        <Typography id="group-slider">Group Size</Typography>
        <Slider
          value={groupCount}
          onChange={(event, newValue) => setGroupCount(newValue)}
          valueLabelDisplay="auto"
          aria-labelledby="group-slider"
          style={{
            marginLeft: 20,
            marginRight: 20,
            flex: 1,
          }}
          min={1}
          max={50}
          marks={[
            { value: 1, label: "1" },
            { value: 50, label: "50" },
          ]}
        />
      </Box>
    );
  };

  const renderDescription = () => {
    return (
      <div className={classes.formSection}>
        <TextField
          label="Description"
          value={description}
          onChange={(event) => setDescription(event.target.value)}
          multiline
          variant="outlined"
          rows={3}
          className={classes.formInput}
          style={{
            width: "100%",
            maxWidth: "600px",
          }}
        />
      </div>
    );
  };

  const renderTags = () => {
    return (
      <div className={classes.formSection}>
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
              required
              variant="outlined"
              label="Tags"
            />
          )}
        />
      </div>
    );
  };

  const onSubmit = async () => {
    if (!validateForm()) {
      setCreatingMeetup(false);
      return;
    }

    setCreatingMeetup(true);
    const { res, resJSON } = await fetcher.createMeetup({
      title,
      time,
      meetupType,
      meetupURL,
      meetupLocation,
      groupCount,
      description,
      tags,
    });
    if (res.ok) {
      clearForm();
      // Use location.replace instead of location.href so user cannot navigate back to create screen
      window.location.replace(`/meetup/${resJSON.id}`);
    }
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h2" component="h1" style={{ textAlign: "center" }}>
        Create your meetup
      </Typography>
      {error && (
        <Typography variant="body1" color="error">
          Please ensure all required fields (marked with *) are filled out.
        </Typography>
      )}
      <Box component="form" display="flex" flexDirection="column">
        {renderNameInput()}
        {renderDateInput()}
        {renderMeetupLocation()}
        {renderGroupSizeInput()}
        {renderDescription()}
        {renderTags()}
        <Box alignSelf="flex-end">
          <Button
            variant="contained"
            color="primary"
            onClick={onSubmit}
            disabled={creatingMeetup}
            style={{
              marginTop: 20,
            }}
          >
            Create Meetup
          </Button>
        </Box>
      </Box>
    </Container>
  );
}

export default CreateMeetup;
