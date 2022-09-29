import { Container } from "@chakra-ui/react";
import { useContext } from "react";
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";

import { ImagesContext } from "../context/app";
import ImageItem from "./ImageItem";

const Images: React.FC = () => {
  const { images } = useContext(ImagesContext);

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
      <ResponsiveMasonry
        columnsCountBreakPoints={{
          350: 1,
          750: 2,
          900: 3,
        }}
      >
        <Masonry gutter="45px">
          {images?.response?.data.map((img) => (
            <ImageItem key={img.id} img={img} />
          ))}
        </Masonry>
      </ResponsiveMasonry>
    </Container>
  );
};

export default Images;
