import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { SWRDevTools } from "swr-devtools";

const colors = {
  brand: {
    100: "#fff",
    700: "#3db46d",
  },
};

const theme = extendTheme({ colors });

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <SWRDevTools>
      <ChakraProvider theme={theme}>
        <App />
      </ChakraProvider>
    </SWRDevTools>
  </React.StrictMode>
);
