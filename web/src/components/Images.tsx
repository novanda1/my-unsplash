import {
  Box,
  Button,
  Image,
  Input,
  InputGroup,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text,
  Container,
} from "@chakra-ui/react";
import { Formik } from "formik";
import { useCallback, useContext, useState } from "react";
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";

import * as Yup from "yup";
import { ImagesContext } from "../context/app";
import { useDelete } from "../hooks/useImage";
import { TImage } from "../lib/api";

const DeleteImageModal: React.FC<{
  id: string;
  isOpen: boolean;
  handleClose: () => void;
}> = ({ id, isOpen, handleClose }) => {
  let pwSchema = Yup.object({
    pw: Yup.string().required().min(6).label("Password"),
  });

  const { handleDelete } = useDelete();

  return (
    <Modal isOpen={isOpen} onClose={handleClose}>
      <ModalOverlay />
      <ModalContent>
        <Formik
          initialValues={{ pw: "" }}
          validationSchema={pwSchema}
          onSubmit={async (_, actions) => {
            actions.setSubmitting(true);
            await handleDelete(id);
            actions.setSubmitting(false);
            handleClose();
          }}
        >
          {({
            errors,
            touched,
            handleBlur,
            handleChange,
            handleSubmit,
            isSubmitting,
          }) => (
            <form onSubmit={handleSubmit}>
              <ModalHeader>Are you sure?</ModalHeader>
              <ModalBody>
                <InputGroup mb="18px">
                  <Box as="label" w="100%">
                    <Text mb="8px" fontSize={14}>
                      Password
                    </Text>
                    <Input
                      type="text"
                      name="pw"
                      placeholder="******************"
                      onChange={handleChange}
                      onBlur={handleBlur}
                    />
                    <Box mt={1}>
                      {errors.pw && touched.pw && (
                        <Text fontSize={14} color="red.600">
                          {errors.pw}
                        </Text>
                      )}
                    </Box>
                  </Box>
                </InputGroup>
              </ModalBody>
              <ModalFooter>
                <Button
                  type="button"
                  variant="ghost"
                  onClick={handleClose}
                  mr={3}
                >
                  Cancel
                </Button>
                <Button
                  type="submit"
                  colorScheme="red"
                  isLoading={isSubmitting}
                >
                  Delete
                </Button>
              </ModalFooter>
            </form>
          )}
        </Formik>
      </ModalContent>
    </Modal>
  );
};

const ImageItem: React.FC<{ img: TImage }> = ({ img }) => {
  const [isOpen, setIsOpen] = useState(false);

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
      <Image src={img.url} />

      <DeleteImageModal id={img.id} handleClose={handleClose} isOpen={isOpen} />
    </Box>
  );
};

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
      <ResponsiveMasonry columnsCountBreakPoints={{ 350: 1, 750: 2, 900: 3 }}>
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
