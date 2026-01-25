// types for frontend
export type BackendPost = {
  id: number;
  created_at: string;
  body: string;
  author: { id: number; name: string };
  comments: {
    id: number;
    body: string;
    author: { id: number; name: string };
  }[];
};

// Mock data
export const mockPosts: BackendPost[] = [
  {
    id: 1,
    created_at: "2026-01-23",
    body: "Welcome to the forum! This is an announcement post.",
    author: { id: 1, name: "Lee Hayoung" },
    comments: [
      { id: 1, body: "Excited to be here!", author: { id: 2, name: "Alice" } },
      { id: 2, body: "Thanks for sharing!", author: { id: 3, name: "Bob" } },
    ],
  },
  {
    id: 2,
    created_at: "2026-01-24",
    body: "I found a bug in the marketplace listing.",
    author: { id: 4, name: "Charlie" },
    comments: [
      { id: 3, body: "Can you explain more?", author: { id: 5, name: "Diana" } },
    ],
  },
  {
    id: 3,
    created_at: "2026-01-25",
    body: "Looking for teammates for OpenJio event.",
    author: { id: 6, name: "Eve" },
    comments: [],
  },
];
