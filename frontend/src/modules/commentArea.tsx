import { useEffect, useState, useRef } from "react";
import "../styles/commentArea.css";

// Импорт иконок
import BoldIcon from "../assets/bold.svg";
import ItalicIcon from "../assets/italic.svg";
import UnderlineIcon from "../assets/undeline.svg";
import { jwtDecode } from "jwt-decode";
import { useNavigate } from "react-router";

const CommentArea = ({
  BookID,
  onCommentSent,
}: {
  BookID: string;
  onCommentSent: () => void;
}) => {
  const navigate = useNavigate();
  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const [UserID, setUserID] = useState<string>();

  const adjustHeight = () => {
    const textarea = textareaRef.current;
    if (textarea) {
      textarea.style.height = "auto";
      textarea.style.height = `${textarea.scrollHeight}px`;
    }
  };

  const sendComment = async () => {
    const textareaBody = textareaRef.current?.value;

    const response = await fetch("/api/comments/", {
      method: "POST",
      body: JSON.stringify({
        userid: UserID,
        bookid: BookID,
        body: textareaBody,
      }),
    });

    if (response.ok) {
      if (textareaRef.current) {
        textareaRef.current.value = "";
        adjustHeight();
      }

      onCommentSent();
    }
  };

  useEffect(() => {
    adjustHeight();

    const id = localStorage.getItem("UserID");

    if (id) {  
      setUserID(id);
    } else {
      navigate("/auth");
    }
  }, []);

  const applyFormat = (prefix: string, suffix: string) => {
    const textarea = textareaRef.current;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const text = textarea.value;

    const before = text.substring(0, start);
    const selected = text.substring(start, end);
    const after = text.substring(end);

    // Формируем новый текст
    const newText = before + prefix + selected + suffix + after;

    // Обновляем значение
    textarea.value = newText;

    // Восстанавливаем фокус и позицию курсора
    textarea.focus();

    const newCursorPos =
      start + prefix.length + (selected.length > 0 ? selected.length : 0);

    textarea.setSelectionRange(newCursorPos, newCursorPos);

    adjustHeight();
  };

  return (
    <div id="CommentArea">
      <textarea
        ref={textareaRef}
        onInput={adjustHeight}
        placeholder="Добавить комментарий..."
      />
      <div>
        <section id="Modifiers_CommentArea">
          <img
            src={BoldIcon}
            alt="Bold"
            onClick={() => applyFormat("**", "**")}
            title="Жирный"
          />
          <img
            src={ItalicIcon}
            alt="Italic"
            onClick={() => applyFormat("*", "*")}
            title="Курсив"
          />
          <img
            src={UnderlineIcon}
            alt="Underline"
            onClick={() => applyFormat("<u>", "</u>")}
            title="Подчеркнутый"
          />
        </section>
        <button onClick={sendComment}>Отправить</button>
      </div>
    </div>
  );
};

export default CommentArea;
