import { useEffect, useState } from "react";

const getUsername = async (id: string) => {
  const response = await fetch(`/api/users/${id}`);
  const userJSON = await response.json();
  return userJSON.name;
};

const parseMarkdown = (text: string): string => {

  let safeText = text
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;');
  
  safeText = safeText
    .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
    .replace(/\*(.*?)\*/g, '<em>$1</em>');
  
  safeText = safeText.replace(/&lt;u&gt;(.*?)&lt;\/u&gt;/g, '<u>$1</u>');

  return safeText;
};

const Comment = ({
  userID,
  commentBody,
}: {
  userID: string;
  commentBody: string;
}) => {
  const [username, setUsername] = useState<string>("Unknown user");
  const [parsedBody, setParsedBody] = useState<string>("");

  useEffect(() => {
    const fetchUsername = async () => {
      const name = await getUsername(userID);
      setUsername(name);
    };
    fetchUsername();
  }, [userID]);

  useEffect(() => {
    setParsedBody(parseMarkdown(commentBody));
  }, [commentBody]);

  return (
    <div id="bodyComment">
      <div id="iconComment"></div>
      <section>
        <h4>{username}</h4>
        {/* TODO: force to use css styles  */}
        <div dangerouslySetInnerHTML={{ __html: parsedBody }} />
      </section>
    </div>
  );
};

export default Comment;