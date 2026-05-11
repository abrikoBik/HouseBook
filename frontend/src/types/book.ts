export default interface BookT {
  author: string;
  cover_type?: string;
  description?: string;
  genre: string;
  id: string;
  isbn: string;
  language?: string;
  page_count?: number;
  publisher: string;
  title: string;
  year?: number;
  coverImage?: string;
}

export default interface CommentT {
  id: string;
  user_id: string;
  book_id: string;
  body: string;
}
