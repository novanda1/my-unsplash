import "../styles/globals.css";
import type { AppProps } from "next/app";
import { ChakraProvider, extendTheme } from "@chakra-ui/react";

const colors = {
  brand: {
    100: "#fff",
    700: "#3db46d",
  },
};

const theme = extendTheme({ colors });

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ChakraProvider theme={theme}>
      <Component {...pageProps} />
    </ChakraProvider>
  );
}

export default MyApp;
