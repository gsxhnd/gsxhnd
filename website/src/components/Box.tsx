// import { motion } from "motion/react";
// import * as motion from "motion/react-client";

export function Box() {
  return <div style={ball} />;
}

const ball = {
  width: 100,
  height: 100,
  backgroundColor: "#dd00ee",
  borderRadius: "50%",
};

// import { Box } from "./Box";
// <Box client:only="react" />
