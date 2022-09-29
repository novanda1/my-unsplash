import { Box, Button, Text } from "@chakra-ui/react";
import Image from "next/image";
import { useState, useCallback } from "react";
import { Blurhash } from "react-blurhash";
import { TImage } from "../lib/api";
import DeleteImageModal from "./DeleteImageModal";

const ImageItem: React.FC<{ img: TImage }> = ({ img }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [opacity, setOpacity] = useState(1);

  const handleClose = useCallback(() => {
    setIsOpen(false);
  }, []);

  const handleOpen = useCallback(() => {
    setIsOpen(true);
  }, []);

  return (
    <Box
      rounded={16}
      overflow="hidden"
      position="relative"
      _hover={{
        ".hover-content": {
          opacity: 1,
          zIndex: 1,
        },
      }}
    >
      <Button
        className="hover-content"
        position="absolute"
        variant="outline"
        colorScheme="red"
        size="xs"
        top={18}
        right={18}
        opacity={0}
        rounded="38px"
        onClick={handleOpen}
      >
        delete
      </Button>
      <Box
        className="hover-content"
        position="absolute"
        left={0}
        right={0}
        bottom={0}
        px="24px"
        py="24px"
        opacity={0}
        bgGradient="linear(to-b, transparent 0%, gray.600 200%)"
      >
        <Text color="white" fontSize={18} fontWeight={700}>
          {img.label}
        </Text>
      </Box>
      <Image
        src={img.url}
        alt=""
        layout="responsive"
        width={img.w}
        height={img.h}
        onLoadingComplete={() => setOpacity(0)}
      />

      <Box position="absolute" inset={0} opacity={opacity}>
        <Blurhash hash={img.hash} width={"100%"} height={"100%"} />
      </Box>

      <DeleteImageModal id={img.id} handleClose={handleClose} isOpen={isOpen} />
    </Box>
  );
};

export default ImageItem;
