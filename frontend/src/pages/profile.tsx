import { useState, useEffect } from "react";
import SideBar from "../modules/sideBar";
import Header from "../modules/header";
import "../styles/profile.css";

interface LibraryAPI {
  book_count: number;
}

interface UserAPI {
  email: string;
  finished_reading_id: string;
  id: string;
  image_uuid: string;
  login: string;
  name: string;
  now_reading_id: string;
  rating: number;
  trade_ready_id: string;
}

interface TradeT {
  id: string;
  owner_id: string;
  buyer_id: string;
  owner_confirmed: boolean;
  buyer_confirmed: boolean;
  completed_at?: {
    valid: boolean;
    time: string;
  } | null;
  created_at: {
    valid: boolean;
    time: string;
  };
}

/* ==========================================================================
   Суб-компонент одной строки сделки (Загружает имена участников обмена)
   ========================================================================== */
interface TradeItemProps {
  trade: TradeT;
  currentUserId: string;
  onDeleteSuccess: (tradeId: string) => void;
  onUpdateSuccess: (updatedTrade: TradeT) => void;
}

const TradeItem = ({
  trade,
  currentUserId,
  onDeleteSuccess,
  onUpdateSuccess,
}: TradeItemProps) => {
  const [partnerName, setPartnerName] = useState<string>("Загрузка...");
  const [isDeleting, setIsDeleting] = useState<boolean>(false);
  const [isCompleting, setIsCompleting] = useState<boolean>(false);
  const isOwner = trade.owner_id === currentUserId;
  const partnerId = isOwner ? trade.buyer_id : trade.owner_id;

  useEffect(() => {
    const fetchPartnerName = async () => {
      try {
        const response = await fetch(`/api/users/${partnerId}`);
        if (response.ok) {
          const data = await response.json();
          setPartnerName(data.name || `@${data.login}`);
        } else {
          setPartnerName("Пользователь не найден");
        }
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (err: any) {
        setPartnerName(err);
      }
    };

    if (partnerId) fetchPartnerName();
  }, [partnerId]);

  const handleCompleteTrade = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (
      !window.confirm("Вы уверены, что хотите принять и завершить эту сделку?")
    )
      return;

    setIsCompleting(true);
    try {
      const response = await fetch(`/api/trades/${trade.id}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ userid: currentUserId }),
      });

      const result = await response.json();
      if (response.ok) {
        onUpdateSuccess(result as TradeT);
        alert("Сделка успешно завершена!");
      } else if (response.status == 500) {
        if (result["error"] == "unknown error") {
          alert("Такая книга уже существует!");
        }
      }
    } catch (err) {
      console.error(err);
    } finally {
      setIsCompleting(false);
    }
  };

  const handleDeleteTrade = async (e: React.MouseEvent) => {
    e.stopPropagation();
    if (
      !window.confirm("Вы уверены, что хотите отменить и удалить эту сделку?")
    )
      return;

    try {
      setIsDeleting(true);
      const response = await fetch(`/api/trades/${trade.id}`, {
        method: "DELETE",
      });
      if (!response.ok) throw new Error("Не удалось удалить сделку");
      onDeleteSuccess(trade.id);
      // eslint-disable-next-line @typescript-eslint/no-explicit-any
    } catch (err: any) {
      console.error("Delete trade error:", err);
      alert(err.message || "Произошла ошибка при удалении сделки");
    } finally {
      setIsDeleting(false);
    }
  };

  const isCompleted = trade.owner_confirmed && trade.buyer_confirmed;
  const statusText = isCompleted ? "Завершена" : "Ожидает подтверждения";
  const statusClass = isCompleted ? "status-completed" : "status-pending";

  return (
    <div
      className={`trade-list-item ${isDeleting ? "trade-item-deleting" : ""}`}
    >
      <div className="trade-main-info">
        <span className="trade-id">Сделка #{trade.id.slice(0, 8)}</span>
        <span className="trade-partner">
          {isOwner ? "Вы отдаете пользователю: " : "Вы берете у пользователя: "}
          <strong>{partnerName}</strong>
        </span>
      </div>
      <div className="trade-meta-container">
        <div className="trade-meta-info">
          <span className={`trade-status ${statusClass}`}>{statusText}</span>
        </div>

        <div className="trade-actions-block">
          {!isCompleted && (
            <button
              className="trade-accept-btn"
              onClick={handleCompleteTrade}
              disabled={isCompleting || isDeleting}
            >
              {isCompleting ? "..." : "Принять"}
            </button>
          )}

          <button
            className="trade-delete-btn"
            onClick={handleDeleteTrade}
            disabled={isDeleting || isCompleting}
            title="Удалить сделку"
          >
            {isDeleting ? "..." : "✕"}
          </button>
        </div>
      </div>
    </div>
  );
};

/* ==========================================================================
   Основной компонент Профиля
   ========================================================================== */
const Profile = () => {
  const [user, setUser] = useState<UserAPI | null>(null);
  const [trades, setTrades] = useState<TradeT[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // НОВЫЕ СТЭЙТЫ: Количество книг в каждой библиотеке
  const [nowReadingCount, setNowReadingCount] = useState<number>(0);
  const [finishedCount, setFinishedCount] = useState<number>(0);
  const [tradeReadyCount, setTradeReadyCount] = useState<number>(0);

  const getUserId = (): string | null => {
    if (typeof window !== "undefined") {
      return localStorage.getItem("UserID");
    }
    return null;
  };

  useEffect(() => {
    const loadProfileAndTrades = async () => {
      try {
        setLoading(true);
        setError(null);

        const userId = getUserId();
        if (!userId) throw new Error("Пользователь не авторизован");

        // Загрузка данных профиля пользователя
        const userResponse = await fetch(`/api/users/${userId}`);
        if (!userResponse.ok) throw new Error("Ошибка загрузки профиля");

        const userData: UserAPI = await userResponse.json();
        setUser(userData);

        // Параллельная загрузка сделок и счетчиков из библиотек
        await Promise.all([
          // Загрузка сделок
          fetch(`/api/users/${userId}/trades`)
            .then((res) => (res.ok ? res.json() : []))
            .then((data) => setTrades(data || []))
            .catch((err) => console.error("Failed to load trades:", err)),

          // Загрузка счетчика "Читает сейчас"
          userData.now_reading_id
            ? fetch(`/api/libraries/${userData.now_reading_id}`)
                .then((res) =>
                  res.ok ? (res.json() as Promise<LibraryAPI>) : null,
                ) // <-- Явно приводим тип
                .then((data) => setNowReadingCount(data?.book_count || 0))
                .catch(() => setNowReadingCount(0))
            : Promise.resolve(setNowReadingCount(0)),

          // Загрузка счетчика "Прочитано"
          userData.finished_reading_id
            ? fetch(`/api/libraries/${userData.finished_reading_id}`)
                .then((res) =>
                  res.ok ? (res.json() as Promise<LibraryAPI>) : null,
                ) // <-- Явно приводим тип
                .then((data) => setFinishedCount(data?.book_count || 0))
                .catch(() => setFinishedCount(0))
            : Promise.resolve(setFinishedCount(0)),

          // Загрузка счетчика "Готов к обмену"
          userData.trade_ready_id
            ? fetch(`/api/libraries/${userData.trade_ready_id}`)
                .then((res) =>
                  res.ok ? (res.json() as Promise<LibraryAPI>) : null,
                ) // <-- Явно приводим тип
                .then((data) => setTradeReadyCount(data?.book_count || 0))
                .catch(() => setTradeReadyCount(0))
            : Promise.resolve(setTradeReadyCount(0)),
        ]);
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } catch (error: any) {
        console.error("Failed to load profile:", error);
        setError(error.message || "Не удалось загрузить профиль");
      } finally {
        setLoading(false);
      }
    };

    loadProfileAndTrades();
  }, []);

  if (loading) {
    return (
      <section className="profile-page">
        <SideBar />
        <main className="profile-content">
          <Header />
          <div style={{ textAlign: "center", padding: "100px" }}>
            Загрузка профиля...
          </div>
        </main>
      </section>
    );
  }

  if (error || !user) {
    return (
      <section className="profile-page">
        <SideBar />
        <main className="profile-content">
          <Header />
          <div style={{ textAlign: "center", padding: "100px", color: "red" }}>
            {error || "Не удалось загрузить профиль"}
          </div>
        </main>
      </section>
    );
  }

  return (
    <section className="profile-page">
      <SideBar />
      <main className="profile-content">
        <Header />
        <div className="profile-container">
          <div className="profile-header">
            <div className="user-info">
              <div className="avatar-placeholder">
                {user.image_uuid ? (
                  <img src={`/api/images/${user.image_uuid}`} alt="Avatar" />
                ) : (
                  "Фото"
                )}
              </div>
              <div className="user-details">
                <h1>{user.name}</h1>
                <p className="username">
                  @{user.login} - {user.id}
                </p>
                <p className="city">📍 Таганрог</p>
                <button className="edit-btn">Редактировать профиль</button>
              </div>
            </div>

            <div className="stats-card">
              <h2>Статистика</h2>
              <div className="stat-item">
                <span className="stat-icon">⭐</span>
                <span>Рейтинг: {user.rating.toFixed(1)}</span>
              </div>

              {/* Выводим посчитанное количество книг вместо ID */}
              <div className="stat-item">
                <span className="stat-icon">📖</span>
                <span>Читает сейчас: {nowReadingCount}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">✅</span>
                <span>Прочитано: {finishedCount}</span>
              </div>
              <div className="stat-item">
                <span className="stat-icon">🔄</span>
                <span>Готов к обмену: {tradeReadyCount}</span>
              </div>
            </div>
          </div>

          <div className="profile-trades-section">
            <h2>Мои обмены / Сделки</h2>
            <div className="trades-list-container">
              {trades.length > 0 ? (
                trades.map((trade) => (
                  <TradeItem
                    key={trade.id}
                    trade={trade}
                    currentUserId={user.id}
                    onDeleteSuccess={(deletedId) => {
                      setTrades((prevTrades) =>
                        prevTrades.filter((t) => t.id !== deletedId),
                      );
                    }}
                    onUpdateSuccess={(updatedTrade) => {
                      setTrades((prevTrades) =>
                        prevTrades.map((t) =>
                          t.id === updatedTrade.id ? updatedTrade : t,
                        ),
                      );
                    }}
                  />
                ))
              ) : (
                <p className="no-trades-message">
                  У вас пока нет активных или завершенных сделок по обмену.
                </p>
              )}
            </div>
          </div>
        </div>
      </main>
    </section>
  );
};

export default Profile;
