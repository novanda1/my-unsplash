import { Box, Button, Container } from "@chakra-ui/react";
import React, { useContext, useEffect } from "react";
import Masonry from "react-masonry-component";
import { useWindowSize } from "usehooks-ts";
import { ImagesContext } from "../context/app";

import { useSearch } from "../hooks/useImage";
import ImageItem from "./ImageItem";

const PAGE_SIZE = 6;

const Images: React.FC = () => {
  const { width } = useWindowSize();
  const { query, setMutate } = useContext(ImagesContext);

  const images = useSearch({ search: query, limit: PAGE_SIZE });
  const { data, error, isValidating, size, mutate } = images;

  const isLoadingInitialData = !data && !error;
  const isLoadingMore =
    isLoadingInitialData ||
    (size > 0 && data && typeof data[size - 1] === "undefined");
  const isEmpty = data?.[0]?.data.length === 0;
  const isReachingEnd =
    isEmpty || (data && data[data.length - 1]?.data.length < PAGE_SIZE);
  const isRefreshing = isValidating && data && data.length === size;

  useEffect(() => {}, [width]);

  useEffect(() => {
    setMutate({ mutate: images.mutate });
  }, [mutate, images.mutate, setMutate]);

  return (
    <Container maxW="container.xl" pb={50}>
      <Box margin="-0.5rem">
        {isLoadingInitialData ? (
          "Loading..."
        ) : isEmpty ? (
          "No image found"
        ) : (
          <>
            {/* @ts-ignore */}
            <Masonry className="gallery" elementType={"ul"}>
              {images.response?.map((page) =>
                page.data?.map((img) => (
                  <li className="imgContainer" key={img.id}>
                    <ImageItem img={img} />
                  </li>
                ))
              )}
            </Masonry>
          </>
        )}
        {!isLoadingInitialData && !isEmpty && (
          <Button
            onClick={() => images.setSize(images.size + 1)}
            isLoading={images.isLoading}
            isDisabled={isReachingEnd}
            display="block"
            mx="auto"
            mt={10}
            mb={10}
          >
            {isLoadingMore
              ? "Loading more images..."
              : isReachingEnd
              ? "No more images"
              : "Load more images"}
          </Button>
        )}
      </Box>
    </Container>
  );
};

export default Images;
