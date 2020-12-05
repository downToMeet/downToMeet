import React, { useEffect, useState } from "react";
import PropTypes from "prop-types";
import {
  Box,
  Card,
  CardActionArea,
  CardContent,
  Chip,
  Typography,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import { useSelector } from "react-redux";
import { useGoogleMaps } from "./LocationPicker";
import * as fetcher from "../../lib/fetch";

/**
 * Meetup information card. Displays main meetup details including:
 * - title
 * - location
 * - time
 * - tags/interests
 *
 * Used to display meetups listed in [Search](#search) and in [Profile](#profile).
 */
function MeetupCard({ title, time, location, id, owner, tags }) {
  const [locationString, setLocationString] = useState(null);
  const [joined, setJoined] = useState(false);
  const userID = useSelector((state) => state.id);
  const { isReady, geocode } = useGoogleMaps([location.coordinates]);

  useEffect(async () => {
    if (location.name) {
      setLocationString(location.name);
    }
    if (!location.name && !location.url && location.coordinates && isReady()) {
      const controller = new AbortController();

      geocode(
        {
          location: {
            lat: location.coordinates.lat,
            lng: location.coordinates.lon,
          },
        },
        (res) => {
          if (controller.signal.aborted) {
            return;
          }
          if (res[0]) {
            setLocationString(res[0].formatted_address);
          }
        }
      );
      return () => {
        controller.abort();
      };
    }
    if (owner === userID) {
      setJoined(true);
    } else {
      const { res, resJSON } = await fetcher.getMeetupAttendees(id);
      if (res.ok && resJSON.attending) {
        if (resJSON.attending.includes(userID)) {
          setJoined(true);
        }
      }
    }
    return undefined;
  }, [location.coordinates, location.name, location.url]);

  const tagList = tags
    ? tags.map((tag) => (
        <Chip label={tag} key={tag} style={{ marginRight: 5 }} />
      ))
    : null;

  const locale = "en-US";
  const eventTimeOptions = {
    hour: "numeric",
    minute: "numeric",
    day: "numeric",
    month: "long",
    year: "numeric",
    timeZoneName: "short",
  };

  return (
    <Card variant="outlined" style={{ margin: "10px 0px" }}>
      <CardActionArea component={Link} to={`/meetup/${id}`}>
        <CardContent>
          <Box display="flex" flexDirection="column">
            <Typography variant="h5" component="h2">
              {title}
            </Typography>
            <Typography color="textSecondary">
              when: {new Date(time).toLocaleString(locale, eventTimeOptions)}
            </Typography>
            {locationString && <Typography>where: {locationString}</Typography>}
            {location.url &&
              (joined ? (
                <Typography>where: {location.url}</Typography>
              ) : (
                <Typography>where: Online</Typography>
              ))}
            <Box display="flex" mt={1}>
              {tagList}
            </Box>
          </Box>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}

MeetupCard.propTypes = {
  /** Meetup title. */
  title: PropTypes.string.isRequired,
  /** Meetup time in string format, e.g. `"2020-12-03T23:22:46.056+08:00"`. */
  time: PropTypes.string.isRequired,
  /** Location of the meetup. For physical locations, only `coordinates` and `name` are defined. For online meetups, only `url` is defined. */
  location: PropTypes.shape({
    /** object containing latitude and longitude, e.g. `{lat: 34.069107615481094, lon: -118.44521328860678}`. */
    coordinates: PropTypes.shape({
      lat: PropTypes.number,
      lon: PropTypes.number,
    }),
    /** Name of the location, e.g. `"UCLA"`. May be empty for unnamed locations. */
    name: PropTypes.string,
    /** Link for online meetups, e.g. `"https://zoom.us/"` */
    url: PropTypes.string,
  }).isRequired,
  /** `id` of the meetup. Parsed from the URL `/meetup/:id`. */
  id: PropTypes.string.isRequired,
  /** User `id` of the meetup owner/organizer. */
  owner: PropTypes.string.isRequired,
  /** Tagged interests of the meetup. */
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default MeetupCard;
