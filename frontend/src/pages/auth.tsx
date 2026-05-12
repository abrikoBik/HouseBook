import Logo from "../assets/LogoWText.svg";
import AuthPhoto from "../assets/AuthPhoto.png";
import "../styles/auth.css";
import { useState } from "react";
import { useNavigate } from "react-router";

interface RegisterForm {
  email: string;
  login: string;
  name: string;
  password: string;
}

interface LoginForm {
  name: string;
  login: string;
  password: string;
}

const Auth = () => {
  const navigate = useNavigate();
  const [active, setActive] = useState<boolean>(false);
  const [visibility, setVisibility] = useState<boolean>(true);
  const [checkPassVisibility, setCheckVisibility] = useState<boolean>(true);
  const [checkPass, setCheckPass] = useState<string>("");

  const [message, setMessage] = useState<string>("");
  const [errorMessage, setErrorMessage] = useState<boolean>(false);

  const [registerForm, setRegisterForm] = useState<RegisterForm>({
    email: "",
    login: "",
    name: "",
    password: "",
  });

  const [loginForm, setLoginForm] = useState<LoginForm>({
    name: window.navigator.userAgent,
    login: "",
    password: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (active) {
      setRegisterForm({
        ...registerForm,
        [e.target.id]: e.target.value,
      });
    } else {
      setLoginForm({
        ...loginForm,
        [e.target.id]: e.target.value,
      });
    }
  };

  const sendRegister = async () => {
    if (registerForm.password == checkPass) {
      const response = await fetch("api/users", {
        method: "POST",
        body: JSON.stringify(registerForm),
      });

      // TODO: error handler
      if (!response.ok) {
        switch (response.status) {
          case 409:
            setMessage("Такой аккаунт уже существует!");
            break;
          default:
            console.log(response.status);
            return;
        }
      }

      setActive(false);
      setErrorMessage(false);
      setMessage("Аккаунт создан!");
    } else {
      setErrorMessage(true);
      setMessage("Пароли не совпадают!");
    }
  };

  const sendLogin = async () => {
    const sessionResponse = await fetch(`/api/sessions`, {
      method: "POST",
      body: JSON.stringify(loginForm),
    });

    // TODO: error handler
    if (!sessionResponse.ok) {
      switch (sessionResponse.status) {
        case 401:
          setErrorMessage(true);
          setMessage("Неверный пароль!");
          break;
        case 404:
          setErrorMessage(true);
          setMessage("Пользователь не найден!");
          break;
        default:
          console.log(sessionResponse.status);
          return;
      }
    }

    const responseJSON = await sessionResponse.json();
    // TODO: install JWT unhash library to get user id
    // TODO: save jwt to cookies
    localStorage.setItem("token", responseJSON["jwt"]);
    navigate("/profile");
  };

  return (
    <section id="bodyAuth">
      <section id="AuthSection">
        <div id="LabelContainer">
          <img src={Logo} alt="logo" id="LabelPhoto" />
          <div>
            <h2>С возвращением!</h2>
            <h5>Мы рады видеть вас</h5>
          </div>
        </div>

        <div id="SliderButton">
          <button
            className={`AuthButtons ${active == false ? "active" : "normal"}`}
            onClick={() => {
              setActive(false);
              setMessage("");
            }}
          >
            Авторизация
          </button>
          <button
            className={`AuthButtons ${active == true ? "active" : "normal"}`}
            onClick={() => {
              setActive(true);
              setMessage("");
            }}
          >
            Регистрация
          </button>
          <div
            className={`selector ${active == true ? "toAuth" : "toRegister"}`}
          ></div>
        </div>

        <div id={`${message == "" ? "HideMessage" : "ShowMessage"}`}>
          <div
            id="MessageWindow"
            className={`${errorMessage ? "ErrorMessage" : ""}`}
          >
            <span
              className="material-symbols-outlined"
              id={`${errorMessage ? "" : "Hidden"}`}
              style={{ marginRight: "0.25rem" }}
            >
              error
            </span>
            {message}
          </div>
        </div>

        <form
          id="Form"
          className={`${active == true ? "FormToAuth" : "FormToRegister"}`}
        >
          {!active ? (
            <>
              <div className="AuthInput">
                <input
                  key="auth-login"
                  id="login"
                  type="text"
                  placeholder="Введите ваш логин"
                  onChange={handleChange}
                />
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  person
                </span>
              </div>
              <div className="AuthInput">
                <input
                  key="auth-password"
                  id="password"
                  type={`${visibility ? "password" : "text"}`}
                  placeholder="Введите ваш пароль"
                  onChange={handleChange}
                />
                <span
                  className="material-symbols-outlined NoSelection VisibilityIcon"
                  onClick={() => setVisibility(!visibility)}
                >
                  {visibility ? "visibility" : "visibility_off"}
                </span>
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  password_2
                </span>
              </div>
              <div id="ForgotPasswordContainer">
                <div>
                  <div>
                    <input type="checkbox" id="AuthCheckBox" />
                    <span className="material-symbols-outlined checkBoxIcon NoSelection">
                      check
                    </span>
                  </div>
                  <label htmlFor="AuthCheckBox" className="NoSelection">
                    Запомнить меня
                  </label>
                </div>
                <a href="/auth" className="NoSelection">
                  Забыли пароль?
                </a>
              </div>
              <button type="button" id="SubmitButton" onClick={sendLogin}>
                Авторизоваться
              </button>
            </>
          ) : (
            <>
              <div className="AuthInput">
                <input
                  key="register-email"
                  id="email"
                  type="text"
                  placeholder="Введите email"
                  onChange={handleChange}
                />
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  mail
                </span>
              </div>
              <div className="AuthInput">
                <input
                  key="register-login"
                  id="login"
                  type="text"
                  placeholder="Введите логин"
                  onChange={handleChange}
                />
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  person
                </span>
              </div>
              <div className="AuthInput">
                <input
                  id="name"
                  type="text"
                  placeholder="Введите ФИО"
                  onChange={handleChange}
                />
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  badge
                </span>
              </div>
              <div className="AuthInput">
                <input
                  id="password"
                  type={`${visibility ? "password" : "text"}`}
                  placeholder="Введите пароль"
                  onChange={handleChange}
                />
                <span
                  className="material-symbols-outlined NoSelection VisibilityIcon"
                  onClick={() => setVisibility(!visibility)}
                >
                  {visibility ? "visibility" : "visibility_off"}
                </span>
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  password_2
                </span>
              </div>
              <div className="AuthInput">
                <input
                  id="checkPass"
                  type={`${checkPassVisibility ? "password" : "text"}`}
                  placeholder="Введите пароль повторно"
                  onChange={(e) => setCheckPass(e.target.value)}
                />
                <span
                  className="material-symbols-outlined NoSelection VisibilityIcon"
                  onClick={() => setCheckVisibility(!checkPassVisibility)}
                >
                  {checkPassVisibility ? "visibility" : "visibility_off"}
                </span>
                <span className="material-symbols-outlined NoSelection PersonIcon">
                  password_2
                </span>
              </div>
              <button type="button" id="SubmitButton" onClick={sendRegister}>
                Зарегестрироваться
              </button>
            </>
          )}
        </form>
      </section>
      <img src={AuthPhoto} alt="authphoto" id="PhotoSection" />
    </section>
  );
};

export default Auth;
