import { useEffect, useRef } from "react";
import "../styles/commentArea.css";

// Импорт иконок
import BoldIcon from "../assets/bold.svg";
import ItalicIcon from "../assets/italic.svg";
import UnderlineIcon from "../assets/undeline.svg";

const CommentArea = () => {
  const textareaRef = useRef<HTMLTextAreaElement>(null);

  const adjustHeight = () => {
    const textarea = textareaRef.current;
    if (textarea) {
      textarea.style.height = "auto";
      textarea.style.height = `${textarea.scrollHeight}px`;
    }
  };

  const sendComment = async () => {
    const book = "019df3c0-2bd4-789f-8e57-c220660e4ff9";
    const user = "019df3c2-50ee-797c-bcce-4c40bf2a00bc";
    const textareaBody = textareaRef.current?.value;

    await fetch("/api/comments/", {
      method: "POST",
      body: JSON.stringify({ userid: user, bookid: book, body: textareaBody }),
    });
  };

  useEffect(() => {
    adjustHeight();
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
