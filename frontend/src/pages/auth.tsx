import Logo from "../assets/LogoWText.svg";
import AuthPhoto from "../assets/AuthPhoto.png";
import "../styles/auth.css";
import { useState } from "react";

const Auth = () => {
  const [active, setActive] = useState<boolean>(false);
  const [visibility, setVisibility] = useState<boolean>(false);

  return (
    <section id="body">
      <section id="AuthSection">
        <div id="LabelContainer">
          <img src={Logo} alt="logo" id="LabelPhoto" />
          <div>
            <h1>С возвращением!</h1>
            <h4>Мы рады видеть вас</h4>
          </div>
        </div>

        <div id="SliderButton">
          <button
            className={`AuthButtons ${active == false ? "active" : "normal"}`}
            onClick={() => setActive(false)}
          >
            Авторизация
          </button>
          <button
            className={`AuthButtons ${active == true ? "active" : "normal"}`}
            onClick={() => setActive(true)}
          >
            Регистрация
          </button>
          <div
            className={`selector ${active == true ? "toAuth" : "toRegister"}`}
          ></div>
        </div>

        {!active ? (
          <form id="Form" action="submit">
            <div className="AuthInput">
              <input type="text" placeholder="Введите ваш логин" />
              <span className="material-symbols-outlined NoSelection PersonIcon">
                person
              </span>
            </div>
            <div className="AuthInput">
              <input
                type={`${visibility ? "password" : "text"}`}
                placeholder="Введите ваш пароль"
              />
              <span
                className="material-symbols-outlined NoSelection VisibilityIcon"
                onClick={() => setVisibility(!visibility)}
              >
                {visibility ? "visibility" : "visibility_off"}
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
            <button type="submit" id="SubmitButton">
              Авторизоваться
            </button>
          </form>
        ) : (
          <form id="Form" action="submit">
            <div className="AuthInput">
              <input type="text" placeholder="Введите email" />
              <span className="material-symbols-outlined NoSelection PersonIcon">
                mail
              </span>
            </div>
            <div className="AuthInput">
              <input type="text" placeholder="Введите логин" />
              <span className="material-symbols-outlined NoSelection PersonIcon">
                person
              </span>
            </div>
            <div className="AuthInput">
              <input type="text" placeholder="Введите ФИО" />
              <span className="material-symbols-outlined NoSelection PersonIcon">
                badge
              </span>
            </div>
            <div className="AuthInput">
              <input
                type={`${visibility ? "password" : "text"}`}
                placeholder="Введите пароль"
              />
              <span
                className="material-symbols-outlined NoSelection VisibilityIcon"
                onClick={() => setVisibility(!visibility)}
              >
                {visibility ? "visibility" : "visibility_off"}
              </span>
            </div>
            <div className="AuthInput">
              <input
                type={`${visibility ? "password" : "text"}`}
                placeholder="Введите пароль повторно"
              />
              <span
                className="material-symbols-outlined NoSelection VisibilityIcon"
                onClick={() => setVisibility(!visibility)}
              >
                {visibility ? "visibility" : "visibility_off"}
              </span>
            </div>
            <button type="submit" id="SubmitButton">
              Авторизоваться
            </button>
          </form>
        )}
      </section>
      <img src={AuthPhoto} alt="authphoto" id="PhotoSection" />
    </section>
  );
};

export default Auth;
