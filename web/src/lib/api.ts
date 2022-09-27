import { EncodedQuery, objectToSearchString } from "serialize-query-params";

export type SaveImageDTO = {
  label: string;
  url: string;
};

export type TImage = {
  id: string;
  label: string;
  url: string;
  createdAt: number;
};

export type ImageResponse<T = TImage> = {
  status: "error" | "success";
  message: string;
  data: T;
};

export type GetImagesDTO = {
  cursor?: string;
  limit?: number;
};

export type SearchImagesDTO = {
  search: string;
} & GetImagesDTO;

export class API {
  private baseurl: string = import.meta.env.VITE_API_URL;

  public async request(endpoint: string, init?: RequestInit) {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    return fetch(this.baseurl + "/api/v1" + endpoint, {...init, headers}).then((r) =>
      r.json()
    );
  }
}

export class ImageAPI extends API {
  private async getImageRequest({
    endpoint,
    queryParam = {},
  }: {
    endpoint?: string;
    queryParam?: object;
  }) {
    let serializedQuery = objectToSearchString(queryParam as EncodedQuery);

    const queryParamUrl = queryParam ? `?${serializedQuery}` : "";

    return this.request(`/images${endpoint + queryParamUrl}`, {
      method: "GET",
    });
  }

  private async postImageRequest<T = ImageResponse>({
    endpoint,
    body = {},
  }: {
    endpoint?: string;
    body?: object;
  }): Promise<T> {
    const endpointUrl = endpoint ? `/${endpoint}` : "";

    return this.request(`/images${endpointUrl}`, {
      method: "POST",
      body: JSON.stringify(body),
    });
  }

  public async saveImage(params: SaveImageDTO): Promise<ImageResponse> {
    return this.postImageRequest({ body: params });
  }

  public async getImage(id: string): Promise<ImageResponse> {
    return this.getImageRequest({ endpoint: id });
  }

  public async getImages(
    param: GetImagesDTO
  ): Promise<ImageResponse<TImage[]>> {
    return this.getImageRequest({ queryParam: param });
  }

  public deleteImage(id: string): Promise<ImageResponse<boolean>> {
    return this.request("/images/" + id, {
      method: "DELETE",
    });
  }

  public search(param: SearchImagesDTO): Promise<ImageResponse<TImage[]>> {
    return this.getImageRequest({ endpoint: "/search", queryParam: param });
  }
}
