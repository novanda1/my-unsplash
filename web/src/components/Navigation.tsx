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
import { AppContext } from "../App";
import AddImageForm from "./AddImageForm";

import Logo from "./Logo";

const Navigation: React.FC = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [query, setQuery] = useState("");

  const { handleChangeData } = useContext(AppContext);

  const handleQueryChange = (e: any) => {
    setQuery(e.target.value);
    handleChangeData(e.target.value);
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
          <HStack spacing="10" justify="space-between">
            <HStack>
              <Logo />
              <InputGroup>
                <InputLeftElement
                  pointerEvents="none"
                  children={<Search2Icon color="gray.300" />}
                />
                <Input
                  type="text"
                  placeholder="Search by name"
                  value={query}
                  onChange={handleQueryChange}
                />
              </InputGroup>
            </HStack>
            <Button colorScheme="green" fontSize={14} onClick={openModal}>
              Add a photo
            </Button>

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
