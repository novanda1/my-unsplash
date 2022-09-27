import {
  Box,
  Button,
  Container,
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
} from "@chakra-ui/react";
import { Formik } from "formik";
import { useCallback, useRef, useState } from "react";
import Masonry, { ResponsiveMasonry } from "react-responsive-masonry";

import * as Yup from "yup";

const dummy = [
  {
    id: "63315cea1870e187f275171d",
    label: "Jeferson Argueta\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664135917329-601bb51aa850?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275171e",
    label: "Ahmed\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664115532297-a2b3b8c2af38?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw5fHx8ZW58MHx8fHw%3D&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275171f",
    label: "Christine Kozak",
    url: "https://images.unsplash.com/photo-1664096219883-7857e422494d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxMnx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751720",
    label: "Marek Piwnicki\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664125010896-7d9b606c0369?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxNXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751721",
    label: "Pawel Czerwinski",
    url: "https://images.unsplash.com/photo-1664111544499-aa9a6c7de16f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwxN3x8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751722",
    label: "Sung Jin Cho",
    url: "https://images.unsplash.com/photo-1664104995040-49d9a268ab7c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyMHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751723",
    label: "Samsung Memory\nMemory storage made for everyone ↗",
    url: "https://images.unsplash.com/photo-1659535880591-78ed91b35158?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwyMXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751724",
    label: "Joel Lee",
    url: "https://images.unsplash.com/photo-1664111601108-993c57ee124b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyNHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751725",
    label: "Steffen Lemmerzahl",
    url: "https://images.unsplash.com/photo-1664112742677-9478378e31e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwyN3x8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751726",
    label: "Windows\nCreate great things with Windows 11 & Microsoft 365 ↗",
    url: "https://images.unsplash.com/photo-1662581872277-0fd0bf3ae8f6?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwzMHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751727",
    label: "Susie Burleson\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664122802538-19fcff974fea?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzM3x8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751728",
    label: "Kellen Riggin\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664123238749-44f8830365e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNnx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751729",
    label: "Ali Drabo",
    url: "https://images.unsplash.com/photo-1663977574293-f65523f4df89?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzOXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172a",
    label: "Jez Timms\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664107014454-3d5758bc32e5?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0Mnx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172b",
    label: "Jeferson Argueta\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664136162748-88b2d5edaea1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0NHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172c",
    label: "Windows\nCreate great things with Windows 11 & Microsoft 365 ↗",
    url: "https://images.unsplash.com/photo-1662581872342-3f8e0145668f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw0OHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172d",
    label: "Marek Piwnicki\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664125010409-8a7f4d82ec07?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0OXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172e",
    label: "Pawel Czerwinski",
    url: "https://images.unsplash.com/photo-1664037109833-5230a4640662?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1MXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f275172f",
    label: "Boxed Water Is Better\nPlant-based. Build a better planet. ↗",
    url: "https://images.unsplash.com/photo-1659482633453-4e51c3e95112?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHw1M3x8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751730",
    label: "Jeferson Argueta\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1663409189430-418c72ff3bad?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1Nnx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751731",
    label: "Julia Blumberg\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664058267614-ccb00305ee75?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw1OXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751732",
    label: "Collin Ross",
    url: "https://images.unsplash.com/photo-1663877254454-5e669a7b27e6?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2Mnx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751733",
    label: "Martin Eriksson",
    url: "https://images.unsplash.com/photo-1664033333006-6f8d1e29bd14?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751734",
    label: "Pawel Czerwinski",
    url: "https://images.unsplash.com/photo-1664040271546-be29247975af?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2N3x8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
  {
    id: "63315cea1870e187f2751735",
    label: "Susan Wilkinson\nAvailable for hire",
    url: "https://images.unsplash.com/photo-1664042913846-abb18eba6226?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MHx8fGVufDB8fHx8&w=1000&q=80",
    createdAt: 1664179434,
  },
];

const DeleteImageModal: React.FC<{
  id: string;
  isOpen: boolean;
  handleClose: () => void;
}> = ({ id, isOpen, handleClose }) => {
  let pwSchema = Yup.object({
    pw: Yup.string().required().min(6).label("Password"),
  });

  return (
    <Modal isOpen={isOpen} onClose={handleClose}>
      <ModalOverlay />
      <ModalContent>
        <Formik
          initialValues={{ pw: "" }}
          validationSchema={pwSchema}
          onSubmit={(values, actions) => {
            console.log("delete image id:", id);

            setTimeout(() => {
              alert(JSON.stringify(values, null, 2));
              actions.setSubmitting(false);
              handleClose();
            }, 1000);
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

const ImageItem: React.FC<{ img: typeof dummy[0] }> = ({ img }) => {
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
  const [data] = useState(dummy);

  return (
    <Container maxW="container.xl" pb={50}>
      <ResponsiveMasonry columnsCountBreakPoints={{ 350: 1, 750: 2, 900: 3 }}>
        <Masonry gutter="45px">
          {data.map((img) => (
            <ImageItem key={img.id} img={img} />
          ))}
        </Masonry>
      </ResponsiveMasonry>
    </Container>
  );
};

export default Images;
