import { createContext } from "react";
import { UseImage } from "../hooks/useImage";
import { ImageResponse, TImage } from "../lib/api";

export interface IImagesContext {
  images: UseImage | undefined;
  handleChangeData: (q: string) => void;
}

export const ImagesContext = createContext<IImagesContext>({
  images: undefined,
  handleChangeData: () => {},
});

export const ImagesContextProvider = ImagesContext.Provider;
