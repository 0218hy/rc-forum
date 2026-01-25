import React, { useEffect, useState } from "react";
import { Box, Typography } from "@mui/material";
import PostCard from "./Postcard";
import { getPosts, type BackendPost } from "../api/posts";
import {mockPosts} from "../api/mockPosts";

export default function Feed() {
  // to show sample for frontend
  const [posts, setPosts] = useState<BackendPost[]>(mockPosts);
  const [loading, setLoading] = useState(true);
  

  useEffect(() => {
    getPosts() // later: getPosts("announcement"), for now get all post
      .then(setPosts)
      .catch(console.error)
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return <Typography align="center">Loading...</Typography>;
  }

  return (
    <Box sx={{ maxWidth: 600, mx: "auto" }}>
      {posts.map((p) => (
        <PostCard
          key={p.id}
          author={p.author.name} 
          createdAt={p.created_at}
          content={p.body}
          comments = {p.comments.map((c) => ({
            id: c.id,
            author: c.author.name,
            content: c.body,
          }))}
        />
      ))}
    </Box>
  );
}
