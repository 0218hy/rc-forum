import React from "react";
import { Box } from "@mui/material";
import PostCard from "./Postcard"

export default function Feed() {
  const posts = [
    {
      id: 1,
      author: "Lee Hayoung",
      createdAt: "2026-01-23",
      content: "This is my first post on the forum!",
      comments: [
        { id: 1, author: "Alice", content: "Nice post!" },
        { id: 2, author: "Bob", content: "Welcome!" },
      ],
    },
    {
      id: 1,
      author: "Lee Hayoung",
      createdAt: "2026-01-23",
      content: "This is my first post on the forum!",
      comments: [
        { id: 1, author: "Alice", content: "Nice post!" },
        { id: 2, author: "Bob", content: "Welcome!" },
      ],
    },
    {
      id: 1,
      author: "Lee Hayoung",
      createdAt: "2026-01-23",
      content: "This is my first post on the forum!",
      comments: [
        { id: 1, author: "Alice", content: "Nice post!" },
        { id: 2, author: "Bob", content: "Welcome!" },
      ],
    },
  ];

  return (
    <Box sx={{ maxWidth: 600, mx: "auto" }}>
      {posts.map((p) => (
        <PostCard
          key={p.id}
          author={p.author}
          createdAt={p.createdAt}
          content={p.content}
          initialComments={p.comments}
        />
      ))}
    </Box>
  );
}
