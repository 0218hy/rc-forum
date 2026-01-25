import React, { useState } from "react";
import {
  Card,
  CardHeader,
  CardContent,
  CardActions,
  Avatar,
  IconButton,
  Typography,
  TextField,
  Button,
  Box,
  Divider,
  Accordion,
  AccordionSummary,
  AccordionDetails,
} from "@mui/material";
import MoreVertIcon from "@mui/icons-material/MoreVert";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";

// Colors
const MAROON = "#800000";
const LIGHT_PINK = "#f9e8eaff";

// Types
type Comment = { id: number; author: string; content: string };
type PostProps = {
    author: string;
    createdAt: string;
    content: string;
    comments: Comment[];
};
  
export default function PostCard({
  author,
  createdAt,
  content,
  comments
}: PostProps) {
  const [newComment, setNewComment] = useState("");

  return (
    <Card sx={{ mb: 3, borderRadius: 3, boxShadow: 3 }}>
      {/* Header */}
      <CardHeader
        avatar={
          <Avatar sx={{ bgcolor: MAROON, color: "#fff" }}>
            {author[0]}
          </Avatar>
        }
        action={
          <IconButton>
            <MoreVertIcon />  
          </IconButton>
        }
        title={
          <Typography align="left" fontWeight="bold" sx={{ color: MAROON }}>
            {author}
          </Typography>
        }
        subheader={
            <Typography align="left">
                {createdAt}
            </Typography>
        }
      />

      {/* Post Content */}
      <CardContent sx={{ backgroundColor: LIGHT_PINK }}>
        <Typography align="left">{content}</Typography>
      </CardContent>

      <Divider />

      {/* Comments */}
      <Accordion
        elevation={0}
        sx={{
          "&:before": { display: "none" },
        }}
      >
        <AccordionSummary expandIcon={<ExpandMoreIcon />}>
          <Typography sx={{ color: MAROON, fontWeight:400 }}>
            {comments.length} Comment{comments.length !== 1 && "s"}
          </Typography>
        </AccordionSummary>

        <AccordionDetails>
          <Box>
            {comments.length === 0 && (
              <Typography
                variant="body2"
                fontStyle="italic"
                sx={{ mb: 1}}
              >
                No comments yet
              </Typography>
            )}

            {comments.map((c, i) => (
              <Box key={c.id}>
                <Box sx={{ py: 1 }}>
                  <Typography align="left" variant="body2">
                    <strong>{c.author}: </strong> {c.content}
                  </Typography>
                </Box>
                {i !== comments.length - 1 && <Divider />}
              </Box>
            ))}

            {/* Add Comment */}
            <CardActions sx={{ px: 0, pt: 1 }}>
              <TextField
                size="small"
                placeholder="Add a comment"
                fullWidth
                value={newComment}
                onChange={(e) => setNewComment(e.target.value)}
                sx={{ mr: 1 }}
              />
              <Button
                size="small"
                variant="contained"
                sx={{
                  bgcolor: MAROON,
                  "&:hover": { bgcolor: "#660000" },
                }}
              >
                Post
              </Button>
            </CardActions>
          </Box>
        </AccordionDetails>
      </Accordion>
    </Card>
  );
}
