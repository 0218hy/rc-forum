export type BackendPost = {
    id: number;
    body: string;
    created_at: string;
    author: {
      id: number;
      name: string;
    };
    comments: {
      id: number;
      body: string;
      author: { id: number; name: string };
    }[];
};

// get post based on type
// if no type, give all posts
export async function getPosts(type?: string): Promise<BackendPost[]> {
    const url = type ? `/api/posts?type=${type}` : "/api/posts"
    const res = await fetch(url);
    if (!res.ok) {
        throw new Error("Failed to fetch posts");
    }
    return res.json();
}