import React from "react";
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

function MeetupCard({ title, time, location, id, tags }) {
  let locationString;
  if (location.name) {
    locationString = location.name;
  } else if (location.url) {
    // TODO: "Online" if not joined
    locationString = location.url;
  }
  // TODO: convert location coords to an address

  const tagList = tags.map((tag) => <Chip label={tag} key={tag} />);

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
  tags: PropTypes.arrayOf(PropTypes.string).isRequired,
};

export default MeetupCard;
