import { Search2Icon } from "@chakra-ui/icons";
import {
  Box,
  Button,
  Container,
  HStack,
  Input,
  InputGroup,
  InputLeftElement,
  Modal,
  ModalContent,
  ModalOverlay,
} from "@chakra-ui/react";
import React, { useCallback, useContext, useState } from "react";
import { ImagesContext } from "../context/app";
import AddImageForm from "./AddImageForm";

import Logo from "./Logo";

const Navigation: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);

  const { setQuery, query } = useContext(ImagesContext);

  const handleQueryChange = (e: any) => {
    setQuery(e.target.value);
  };

  const openModal = useCallback(() => {
    setIsOpen(true);
  }, []);

  const closeModal = useCallback(() => {
    setIsOpen(false);
  }, []);

  return (
    <Box as="section" pb={{ base: "12", md: "51px" }}>
      <Box as="nav" bg="bg-surface">
        <Container py={{ base: "4", lg: "5" }} maxW="container.xl">
          <HStack
            spacing="10"
            justify={["flex-start", "space-between"]}
            flexWrap={["wrap", "nowrap"]}
            gap={[5]}
          >
            <Logo />

            <HStack
              justifyContent="space-between"
              w="100%"
              style={{ marginLeft: 0 }}
            >
              <InputGroup w="max-content" flex={[1, "unset"]}>
                <InputLeftElement pointerEvents="none">
                  <Search2Icon color="gray.300" />
                </InputLeftElement>
                <Input
                  type="text"
                  placeholder="Search by name"
                  value={query}
                  onChange={handleQueryChange}
                />
              </InputGroup>
              <Button colorScheme="green" fontSize={14} onClick={openModal}>
                Add a photo
              </Button>
            </HStack>

            <Modal isOpen={isOpen} onClose={closeModal}>
              <ModalOverlay />
              <ModalContent>
                <Box py={"24px"} px={"32px"}>
                  <AddImageForm handleClose={closeModal} />
                </Box>
              </ModalContent>
            </Modal>
          </HStack>
        </Container>
      </Box>
    </Box>
  );
};

export default Navigation;
