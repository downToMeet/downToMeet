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
                <Typography>
                  where:{" "}
                  <Typography component="a" href={location.url}>
                    {location.url}
                  </Typography>
                </Typography>
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
  title: PropTypes.string.isRequired,
  time: PropTypes.string.isRequired,
  location: PropTypes.shape({
    coordinates: PropTypes.shape({
      lat: PropTypes.number,
      lon: PropTypes.number,
    }),
    name: PropTypes.string,
    url: PropTypes.string,
  }).isRequired,
  id: PropTypes.string.isRequired,
  owner: PropTypes.string.isRequired,
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default MeetupCard;
