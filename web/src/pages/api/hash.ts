// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type {
  NextApiRequest,
  NextApiResponse,
} from "next";
import {
  getPlaiceholder,
  IGetPlaiceholderReturn,
} from "plaiceholder";

import fetch from "node-fetch";

type Data = Partial<IGetPlaiceholderReturn>;

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  if (req.method?.toLowerCase() !== "post") {
    res.status(400).json({});
  }

  const image = await fetch(
    req.body.url as string
  );

  const imageBuffer = Buffer.from(
    await image.arrayBuffer()
  );

  const plaice = await getPlaiceholder(
    imageBuffer
  ).then((img) => {
    return img;
  });

  res.status(200).json(plaice);
}
