import { Container } from "@chakra-ui/react";
import React, { useContext, useEffect } from "react";
import Masonry from "react-masonry-component";
import { useWindowSize } from "usehooks-ts";

import { ImagesContext } from "../context/app";
import ImageItem from "./ImageItem";

const Images: React.FC = () => {
  const { images } = useContext(ImagesContext);
  const { width } = useWindowSize();

  useEffect(() => {}, [width]);

  if (images?.isLoading)
    return (
      <Container maxW="container.xl" pb={50}>
        Loading...
      </Container>
    );

  if (images?.isError)
    return (
      <Container maxW="container.xl" pb={50}>
        Something went wrong...
      </Container>
    );

  if (!images?.isError && !images?.isLoading && !images?.response?.data.length)
    return (
      <Container maxW="container.xl" pb={50}>
        No result...
      </Container>
    );

  return (
    <Container maxW="container.xl" pb={50}>
      {/* @ts-ignore */}
      <Masonry className="gallery" elementType={"ul"}>
        {images?.response?.data.map((img) => (
          <li className="imgContainer" key={img.id}>
            <ImageItem img={img} />
          </li>
        ))}
      </Masonry>
    </Container>
  );
};

export default Images;
