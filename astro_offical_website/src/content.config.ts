import { defineCollection, z } from "astro:content";
import { glob } from "astro/loaders";

const newsCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "./src/content/news" }),
  schema: z.object({
    title: z.string(),
    date: z.date(),
    type: z.string().optional(),
    description: z.string().optional(),
    draft: z.boolean().optional().default(false),
  }),
});

export const collections = {
  news: newsCollection,
};
