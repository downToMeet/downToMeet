import React, { useState, useEffect } from "react";
import PropTypes from "prop-types";
import { useSelector } from "react-redux";
import { useLocation, useHistory } from "react-router-dom";
import makeStyles from "@material-ui/styles/makeStyles";
import {
  Box,
  Button,
  Container,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  FormControl,
  InputLabel,
  Select,
  Slider,
  TextField,
  Typography,
} from "@material-ui/core";
import {
  KeyboardDateTimePicker,
  MuiPickersUtilsProvider,
} from "@material-ui/pickers";
import DayUtils from "@date-io/dayjs";

import LocationPicker from "../common/LocationPicker";
import TagPicker from "../common/TagPicker";
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

/**
 * Page for creating or editing meetups, shown on paths `/create` and `/meetup/:id/edit`.
 * The create meetup view is shown on path `/create`, when `:id` is not given.
 * The edit view is shown on path `/meetup/:id/edit`, if `:id` is the ID of an existing meetup
 * and the user is the owner of that meetup. If `id` is provided but the user is not the owner,
 * they are redirected to the [info page](#meetup) for that meetup. If `id` does not match an
 * existing meetup, the user is redirected to `/create`. Note: changes to attendee capacity after
 * meetup creation is not allowed.
 *
 * If info for the current user is not found, they are shown the create view with form submission
 * disabled. If the user is logged in, the correct view will be shown once their data is loaded.
 */
function CreateMeetup({ id }) {
  const location = useLocation(); // listed as dependency in useEffect(). refreshes component if user clicks New Meetup while on edit page.
  const history = useHistory();

  const classes = useStyles();
  const userID = useSelector((state) => state.id);
  const [isEdit, setIsEdit] = useState(false);
  const [isCancelled, setIsCancelled] = useState(false);
  const [title, setTitle] = useState("");
  const [time, setTime] = useState(new Date());
  const [meetupType, setMeetupType] = useState("");
  const [meetupURL, setMeetupURL] = useState("");
  const [meetupLocation, setMeetupLocation] = useState(null);
  const [groupCount, setGroupCount] = useState([2, 10]);
  const [description, setDescription] = useState("");
  const [tags, setTags] = useState([]);

  const [error, setError] = useState(false);
  const [showConfirm, setShowConfirm] = useState(false);
  const [updatingMeetup, setUpdatingMeetup] = useState(false);

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
    setUpdatingMeetup(false);
  };

  useEffect(async () => {
    // TODO: redirect to login page for unauthenticated users.
    // - Redux store may need a 'logging in' action to distinguish between unauthenticated
    //   users and authenticated users for whom userID is not loaded in store yet.
    if (!id || !userID) {
      return;
    }
    const { res, resJSON } = await fetcher.getMeetup(id);
    // if meetup does not exist (or error fetching meetup), redirect to /create
    if (!res.ok) {
      history.replace(`/create`);
      return;
    }
    // if user is not meetup owner, redirect to meetup info page /meetup/:id
    if (userID !== resJSON.owner) {
      history.replace(`/meetup/${id}`);
      return;
    }
    setIsEdit(true);
    setIsCancelled(Boolean(resJSON.canceled));
    setTitle(resJSON.title);
    setTime(new Date(resJSON.time));
    setMeetupType(resJSON.location.url ? REMOTE : IN_PERSON);
    if (resJSON.location.coordinates) {
      setMeetupLocation({
        description: resJSON.location.name,
        coords: [
          resJSON.location.coordinates.lat,
          resJSON.location.coordinates.lon,
        ],
      });
    }
    setMeetupURL(resJSON.location.url || "");
    setGroupCount([resJSON.minCapacity, resJSON.maxCapacity]);
    setDescription(resJSON.description);
    setTags(resJSON.tags);
  }, [userID, location]);

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
          disabled={isCancelled}
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
            disabled={isCancelled}
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
            disabled={isCancelled}
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
            disabled={isCancelled}
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
            disabled={isCancelled}
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
          disabled={isEdit || isCancelled}
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
          disabled={isCancelled}
          value={description}
          onChange={(event) => setDescription(event.target.value)}
          multiline
          variant="outlined"
          rows={3}
          rowsMax={15}
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
        <TagPicker tags={tags} setTags={setTags} />
      </div>
    );
  };

  const handleCancelMeetup = async () => {
    setIsCancelled(true);
    setUpdatingMeetup(true);
    const res = await fetcher.cancelMeetup(id);
    if (res.ok) {
      clearForm();
      history.replace(`/meetup/${id}`);
    } else {
      setUpdatingMeetup(false);
    }
  };

  const onSubmit = async () => {
    if (!validateForm()) {
      setUpdatingMeetup(false);
      return;
    }

    setUpdatingMeetup(true);
    const meetup = {
      id,
      title,
      time,
      meetupType,
      meetupURL,
      meetupLocation,
      groupCount,
      description,
      tags,
    };
    const { res, resJSON } = await fetcher.createOrEditMeetup(meetup, isEdit);
    if (res.ok) {
      clearForm();
      // Use location.replace instead of location.href so user cannot navigate back to create screen
      history.replace(`/meetup/${resJSON.id}`);
    }
  };

  const renderEditButtons = () => {
    const editButtons = (
      <>
        <Button
          variant="contained"
          onClick={() => {
            history.replace(`/meetup/${id}`);
          }}
          disabled={updatingMeetup}
          style={{
            marginTop: 20,
            marginRight: 10,
          }}
        >
          Discard Changes
        </Button>
        <Button
          variant="contained"
          color="secondary"
          disabled={updatingMeetup}
          style={{
            marginTop: 20,
            marginRight: 10,
          }}
          onClick={() => setShowConfirm(true)}
        >
          Cancel Meetup
        </Button>
        <Dialog open={showConfirm} onClose={() => setShowConfirm(false)}>
          <DialogContent>
            <DialogContentText>
              Are you sure you want to cancel the meetup? This cannot be undone.
            </DialogContentText>
            <DialogActions>
              <Button onClick={() => setShowConfirm(false)}>No</Button>
              <Button
                disabled={updatingMeetup}
                onClick={() => {
                  setShowConfirm(false);
                  handleCancelMeetup(id);
                }}
              >
                Yes
              </Button>
            </DialogActions>
          </DialogContent>
        </Dialog>
      </>
    );
    return (
      <Box alignSelf="flex-end">
        {isEdit && !isCancelled && editButtons}
        <Button
          variant="contained"
          color="primary"
          onClick={
            !isCancelled ? onSubmit : () => history.replace(`/meetup/${id}`)
          }
          disabled={!userID || updatingMeetup}
          style={{
            marginTop: 20,
          }}
        >
          {isCancelled && "Back"}
          {!isCancelled && (isEdit ? "Save Changes" : "Create Meetup")}
        </Button>
      </Box>
    );
  };

  return (
    <Container maxWidth="sm">
      <Typography variant="h2" component="h1" style={{ textAlign: "center" }}>
        {isEdit ? "Edit" : "Create"} your meetup
      </Typography>
      {error && (
        <Typography variant="body1" color="error">
          Please ensure all required fields (marked with *) are filled out.
        </Typography>
      )}
      {isCancelled && (
        <Typography variant="body1" color="error">
          This meetup has been cancelled. You can no longer edit the meetup.
        </Typography>
      )}
      {!userID && (
        <Typography variant="body1" color="error">
          You must be logged in to create a meetup.
        </Typography>
      )}
      <Box component="form" display="flex" flexDirection="column">
        {renderNameInput()}
        {renderDateInput()}
        {renderMeetupLocation()}
        {renderGroupSizeInput()}
        {renderDescription()}
        {renderTags()}
        {renderEditButtons()}
      </Box>
    </Container>
  );
}

CreateMeetup.propTypes = {
  /**
   * `id` of the meetup, parsed from the route `/meetup/:id/edit`.
   * Not defined for route `/create`.
   */
  id: PropTypes.string,
};

CreateMeetup.defaultProps = {
  id: "",
};

export default CreateMeetup;
