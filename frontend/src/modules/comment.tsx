import { useEffect, useState } from "react";

const getUsername = async (id: string) => {
  const response = await fetch(`/api/users/${id}`);
  const userJSON = await response.json();
  return userJSON.name;
};

const Comment = ({
  userID,
  commentBody,
}: {
  userID: string;
  commentBody: string;
}) => {
  const [username, setUsername] = useState<string>("Unknown user");

  useEffect(() => {
    const fetchUsername = async () => {
      const name = await getUsername(userID);
      setUsername(name);
    };
    fetchUsername();
  }, [userID]);

  return (
    <div id="bodyComment">
      <div id="iconComment"></div>
      <section>
        <h4>{username}</h4>
        <p>{commentBody}</p>
      </section>
    </div>
  );
};

export default Comment;
