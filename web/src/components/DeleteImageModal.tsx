import {
  Box,
  Button,
  Input,
  InputGroup,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Text
} from "@chakra-ui/react";
import { Formik } from "formik";
import { useDelete } from "../hooks/useImage";

import * as Yup from "yup";

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
                      type="password"
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

export default DeleteImageModal;
