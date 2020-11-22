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
  TextField,
  Typography,
} from "@material-ui/core";
import AutoComplete from "@material-ui/lab/Autocomplete";
import {
  KeyboardDatePicker,
  KeyboardTimePicker,
  MuiPickersUtilsProvider,
} from "@material-ui/pickers";
import DayUtils from "@date-io/dayjs";

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
  const [meetupLocation, setMeetupLocation] = useState("");
  const [groupMin, setGroupMin] = useState(0);
  const [groupMax, setGroupMax] = useState(0);
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
          label="Meetup Name"
          value={name}
          onChange={(event) => setName(event.target.value)}
          className={classes.formInput}
        />
      </div>
    );
  };

  const renderDateInput = () => {
    return (
      <div className={classes.formSection}>
        <MuiPickersUtilsProvider utils={DayUtils}>
          <KeyboardDatePicker
            required
            variant="inline"
            inputVariant="outlined"
            format="MM/DD/YYYY"
            label="Meetup Date"
            value={date}
            onChange={(newDate) => setDate(newDate)}
            className={classes.formInput}
            style={{
              marginRight: 20,
            }}
          />
          <KeyboardTimePicker
            required
            variant="inline"
            inputVariant="outlined"
            label="Meetup Time"
            value={date}
            onChange={(newDate) => setDate(newDate)}
            className={classes.formInput}
          />
        </MuiPickersUtilsProvider>
      </div>
    );
  };

  const renderMeetupLocation = () => {
    return (
      <Box display="flex" flexWrap="wrap" className={classes.formSection}>
        <FormControl required variant="outlined" className={classes.formInput}>
          <InputLabel id="select-meetup-type-label">Meetup Type</InputLabel>
          <Select
            label="Meetup Type"
            labelId="select-meetup-type-label"
            value={meetupType}
            onChange={(event) => setMeetupType(event.target.value)}
          >
            <MenuItem value="in-person">In person</MenuItem>
            <MenuItem value="remote">Remote</MenuItem>
          </Select>
        </FormControl>
        {meetupType === "remote" && ( // TODO: add location support for in-person events
          <TextField
            required
            label="Meetup Location"
            variant="outlined"
            value={meetupLocation}
            onChange={(event) => setMeetupLocation(event.target.value)}
            className={classes.formInput}
            style={{
              marginLeft: 20,
            }}
          />
        )}
      </Box>
    );
  };

  const renderGroupSizeInput = () => {
    return (
      <Box display="flex" alignItems="baseline" className={classes.formSection}>
        <Typography>Group Size: </Typography>
        <TextField
          variant="outlined"
          label="Min"
          type="number"
          helperText="0 for no minimum"
          size="small"
          value={groupMin}
          onChange={(event) => setGroupMin(parseInt(event.target.value, 10))}
          style={{
            margin: "0px 20px",
            width: 100,
          }}
        />
        <Typography> to </Typography>
        <TextField
          variant="outlined"
          label="Max"
          type="number"
          helperText="0 for no maximum"
          size="small"
          value={groupMax}
          onChange={(event) => setGroupMax(parseInt(event.target.value, 10))}
          style={{
            margin: "0px 20px",
            width: 100,
          }}
        />
      </Box>
    );
  };

  const renderDescription = () => {
    return (
      <div className={classes.formSection}>
        <TextField
          label="Meetup Description"
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
              label="Meetup tags"
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
            variant="outlined"
            onClick={onSubmit}
            style={{
              margin: 20,
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
