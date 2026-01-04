import { defineCollection, z } from "astro:content";

const newsCollection = defineCollection({
  type: "content",
  // loader: glob({ pattern: "**/*.{md,mdx}", base: "./src/blog" }),
  schema: z.object({
    title: z.string(),
    // date: z.string(),
    date: z.date(),
    type: z.string().optional(),
    description: z.string().optional(),
    draft: z.boolean().optional().default(false),
  }),
});

export const collections = {
  news: newsCollection,
};
