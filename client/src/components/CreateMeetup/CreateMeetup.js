import { React, useState } from "react";
import makeStyles from "@material-ui/styles/makeStyles";
import {
  Box,
  Button,
  Container,
  FormControl,
  InputLabel,
  MenuItem,
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
  const [name, setName] = useState("");
  const [date, setDate] = useState(new Date());
  const [meetupType, setMeetupType] = useState("");
  const [meetupURL, setMeetupURL] = useState("");
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([2, 10]);
  const [desc, setDesc] = useState("");
  const [tags, setTags] = useState([]);

  // TODO: load options from the server instead
  const tagOptions = [
    "basketball",
    "ping pong",
    "badminton",
    "movie",
    "cooking",
  ];

  const renderNameInput = () => {
    return (
      <div className={classes.formSection}>
        <TextField
          required
          variant="outlined"
          label="Title"
          value={name}
          onChange={(event) => setName(event.target.value)}
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
            value={date}
            onChange={(newDate) => setDate(newDate)}
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
            labelId="select-meetup-type-label"
            value={meetupType}
            onChange={(event) => setMeetupType(event.target.value)}
          >
            <MenuItem value="in-person">In person</MenuItem>
            <MenuItem value="remote">Remote</MenuItem>
          </Select>
        </FormControl>
        {meetupType === "in-person" && (
          <LocationPicker
            value={meetupLocation}
            setValue={setMeetupLocation}
            style={{
              marginLeft: 20,
              flex: 1,
            }}
          />
        )}
        {meetupType === "remote" && (
          <TextField
            label="URL"
            variant="outlined"
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
          value={desc}
          onChange={(event) => setDesc(event.target.value)}
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

  const onSubmit = () => {
    // TODO: disable button while waiting
    // TODO: validate form
    // console.log("Creating meetup: ");
    // console.log({ name, date, meetupType, meetupLocation, groupMin, groupMax, desc, tags });
    // TODO: create a meetup and redirect user
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

  return (
    <Container maxWidth="sm">
      <Typography variant="h2" component="h1">
        Create your meetup
      </Typography>
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
