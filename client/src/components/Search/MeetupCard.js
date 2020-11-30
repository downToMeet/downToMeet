import React from "react";
import PropTypes from "prop-types";
import { Box, Card, CardContent, Chip, Typography } from "@material-ui/core";
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

  return (
    <Card variant="outlined">
      <CardContent>
        <Typography>{title}</Typography>
        <Typography>{new Date(time).toLocaleString()}</Typography>
        <Typography>{locationString}</Typography>
        <Link to={`/meetup/${id}`}>See Details</Link>
        <Box display="flex">{tagList}</Box>
      </CardContent>
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
