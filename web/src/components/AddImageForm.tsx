import {
  Heading,
  InputGroup,
  Input,
  HStack,
  Button,
  Box,
  Text,
} from "@chakra-ui/react";
import { Formik } from "formik";

import * as Yup from "yup";
import { useSaveImage } from "../hooks/useImage";

const AddImageForm: React.FC<{ handleClose: () => void }> = ({
  handleClose,
}) => {
  const imageSchema = Yup.object({
    label: Yup.string().min(2).required().label("Photo Label"),
    url: Yup.string().required().url().label("Photo URL"),
  });

  const { handleSave } = useSaveImage();

  return (
    <Box>
      <Heading as="h3" mb="20px" fontWeight={500} fontSize={24}>
        Add a new photo
      </Heading>
      <Formik
        initialValues={{ url: "", label: "" }}
        validationSchema={imageSchema}
        onSubmit={async (values, actions) => {
          actions.setSubmitting(true);
          await handleSave(values);
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
            <InputGroup mb="18px">
              <Box as="label" w="100%">
                <Text mb="8px" fontSize={14}>
                  Label
                </Text>
                <Input
                  type="text"
                  name="label"
                  placeholder="Rainy day"
                  onChange={handleChange}
                  onBlur={handleBlur}
                />
                <Box mt={1}>
                  {errors.label && touched.label && (
                    <Text fontSize={14} color="red.600">
                      {errors.label}
                    </Text>
                  )}
                </Box>
              </Box>
            </InputGroup>

            <InputGroup mb="18px">
              <Box as="label" w="100%">
                <Text mb="8px" fontSize={14}>
                  Photo URL
                </Text>
                <Input
                  type="text"
                  name="url"
                  placeholder="https://images.unsplash.com/photo-1584395630827-860eee694d7b?ixlib=r..."
                  onChange={handleChange}
                  onBlur={handleBlur}
                />
                <Box mt={1}>
                  {errors.url && touched.url && (
                    <Text fontSize={14} color="red.600">
                      {errors.url}
                    </Text>
                  )}
                </Box>
              </Box>
            </InputGroup>

            <HStack justifyContent="flex-end">
              <Button type="button" variant="ghost" onClick={handleClose}>
                Cancel
              </Button>
              <Button
                type="submit"
                colorScheme="green"
                isLoading={isSubmitting}
              >
                Submit
              </Button>
            </HStack>
          </form>
        )}
      </Formik>
    </Box>
  );
};

export default AddImageForm;
