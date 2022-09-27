import useSWR from "swr";
import { GetImagesDTO, ImageAPI, SearchImagesDTO } from "../lib/api";

const api = new ImageAPI();

export const useImages = (key: GetImagesDTO) => {
  const { data, error } = useSWR(key, (...args) => api.getImages(...args));

  return {
    images: data,
    isLoading: !error && !data,
    isError: error,
  };
};

export const useSearch = (key: SearchImagesDTO) => {
  const { data, error } = useSWR(key, (...args) => api.search(...args));

  return {
    images: data,
    isLoading: !error && !data,
    isError: error,
  };
};
