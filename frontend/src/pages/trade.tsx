import React, { useState } from 'react';
import '../styles/trade.css';
import SideBar from '../modules/sideBar';
import Header from '../modules/header';

interface Book {
  id: number;
  title: string;
  author: string;
  color: string;
}

const myBooks: Book[] = [
  { id: 1, title: 'OKP', author: 'Джонатан Абрамович', color: '#e8e4d4' },
  { id: 2, title: 'Лао-цзы', author: 'Бронислав Виногродский', color: '#e63946' },
  { id: 3, title: 'Духовный путь', author: 'Грегори Дэвид Робертс', color: '#457b9d' },
  { id: 4, title: 'Портрет Дориана Грея', author: 'Оскар Уайльд', color: '#6b21a8' },
  { id: 5, title: 'Облако', author: 'Константин Аксаков', color: '#87ceeb' },
  { id: 6, title: '1984', author: 'Джордж Оруэлл', color: '#555' },
  { id: 7, title: 'Мастер и Маргарита', author: 'Булгаков', color: '#8b4513' },
  { id: 8, title: 'Преступление и наказание', author: 'Достоевский', color: '#4a4a4a' },
];

const partnerBooks: Book[] = [
  { id: 101, title: 'Война и мир', author: 'Лев Толстой', color: '#f4e4d4' },
  { id: 102, title: 'Анна Каренина', author: 'Лев Толстой', color: '#c9a96e' },
  { id: 103, title: 'Евгений Онегин', author: 'Пушкин', color: '#2f4f4f' },
  { id: 104, title: 'Идиот', author: 'Достоевский', color: '#6b4423' },
  { id: 105, title: 'Отцы и дети', author: 'Тургенев', color: '#4682b4' },
  { id: 106, title: 'Герой нашего времени', author: 'Лермонтов', color: '#708090' },
  { id: 107, title: 'Мёртвые души', author: 'Гоголь', color: '#3c1414' },
  { id: 108, title: 'Обломов', author: 'Гончаров', color: '#5f6a6a' },
];

const TradePage: React.FC = () => {
  const [myInventory, setMyInventory] = useState<Book[]>(myBooks);
  const [partnerInventory, setPartnerInventory] = useState<Book[]>(partnerBooks);
  const [mySelected, setMySelected] = useState<Book[]>([]);
  const [partnerSelected, setPartnerSelected] = useState<Book[]>([]);
  const [activeTab, setActiveTab] = useState<'my' | 'partner'>('my');
  const [message, setMessage] = useState('');
  const [searchQuery, setSearchQuery] = useState('');
  const [toast, setToast] = useState('');

  const pickMy = (b: Book) => {
    setMySelected([...mySelected, b]);
    setMyInventory(myInventory.filter(x => x.id !== b.id));
  };
  const pickPartner = (b: Book) => {
    setPartnerSelected([...partnerSelected, b]);
    setPartnerInventory(partnerInventory.filter(x => x.id !== b.id));
  };
  const removeMy = (b: Book) => {
    setMySelected(mySelected.filter(x => x.id !== b.id));
    setMyInventory([...myInventory, b]);
  };
  const removePartner = (b: Book) => {
    setPartnerSelected(partnerSelected.filter(x => x.id !== b.id));
    setPartnerInventory([...partnerInventory, b]);
  };

  const sendOffer = () => {
    if (mySelected.length === 0) {
      setToast('Добавьте хотя бы одну книгу');
      setTimeout(() => setToast(''), 3000);
      return;
    }
    setToast('Предложение отправлено!');
    setTimeout(() => setToast(''), 3000);
  };

  const filteredMy = myInventory.filter(b =>
    b.title.toLowerCase().includes(searchQuery.toLowerCase()) || b.author.toLowerCase().includes(searchQuery.toLowerCase())
  );
  const filteredPartner = partnerInventory.filter(b =>
    b.title.toLowerCase().includes(searchQuery.toLowerCase()) || b.author.toLowerCase().includes(searchQuery.toLowerCase())
  );

  const BookCard = ({ book, onClick, small = false }: { book: Book; onClick: () => void; small?: boolean }) => (
    <div className={`card ${small ? 'sm' : ''}`} onClick={onClick}>
      <div className="card-cover" style={{ background: book.color }}>
        {book.title}
      </div>
      {!small && <div className="card-info"><b>{book.title}</b><span>{book.author}</span></div>}
    </div>
  );

  const EmptyGrid = () => (
    <div className="grid-4x2">
      {Array.from({ length: 8 }).map((_, i) => <div key={i} className="slot" />)}
    </div>
  );

  const Inventory = ({ books, onPick, label, count }: { books: Book[]; onPick: (b: Book) => void; label: string; count: number }) => (
    <div className="inv-section">
      <div className="inv-header">
        <div className="inv-dropdown"><span>📚</span><span>{label}</span><span>▼</span></div>
        <div className="inv-search">
          <input placeholder="Поиск книги..." value={searchQuery} onChange={e => setSearchQuery(e.target.value)} />
        </div>
      </div>
      <div className="inv-grid">
        {books.map(b => <BookCard key={b.id} book={b} onClick={() => onPick(b)} />)}
      </div>
      <div className="inv-footer">
        <span>Показано {books.length} из {count}</span>
        <div className="pagination"><button disabled>←</button><span>1 из 1</span><button disabled>→</button></div>
      </div>
    </div>
  );

  return (
    <div className="trade-page">
      <SideBar />
      <Header />
      
      <div className="trade-wrap">
        {/* LEFT */}
        <div className="panel-left">
          <div className="tabs">
            <button className={`tab ${activeTab === 'my' ? 'active' : ''}`} onClick={() => setActiveTab('my')}>Ваша коллекция</button>
            <button className={`tab ${activeTab === 'partner' ? 'active' : ''}`} onClick={() => setActiveTab('partner')}>Коллекция партнёра</button>
          </div>
          {activeTab === 'my' ? (
            <Inventory books={filteredMy} onPick={pickMy} label="Моя библиотека" count={myBooks.length} />
          ) : (
            <Inventory books={filteredPartner} onPick={pickPartner} label="Коллекция: Алексей Петров" count={partnerBooks.length} />
          )}
        </div>

        {/* RIGHT */}
        <div className="panel-right">
          {/* My selected */}
          <div className="summary-box">
            <div className="summary-head">
              <span>📚</span>
              <div><b>Ваши книги:</b><small>Книги, которые вы отдадите</small></div>
            </div>
            {mySelected.length > 0 ? (
              <div className="grid-4x2">
                {mySelected.map(b => <BookCard key={b.id} book={b} onClick={() => removeMy(b)} small />)}
              </div>
            ) : <EmptyGrid />}
            {mySelected.length === 0 && <p className="hint">Ожидание выбора одной или нескольких книг</p>}
          </div>

          {/* Arrows */}
          <div className="arrows"><span>↗</span><span>↙</span></div>

          {/* Partner selected */}
          <div className="summary-box">
            <div className="summary-head">
              <span>👤</span>
              <div><b>Книги Алексея:</b><small>Книги, которые вы получите</small></div>
            </div>
            {partnerSelected.length > 0 ? (
              <div className="grid-4x2">
                {partnerSelected.map(b => <BookCard key={b.id} book={b} onClick={() => removePartner(b)} small />)}
              </div>
            ) : <EmptyGrid />}
          </div>

          <p className="notify">Алексей будет уведомлён об этом предложении</p>

          <div className="msg-box">
            <label>Сообщение к предложению:</label>
            <textarea placeholder="Напишите сообщение..." value={message} onChange={e => setMessage(e.target.value)} rows={3} />
          </div>

          <button className="btn-primary" onClick={sendOffer} disabled={mySelected.length === 0}>Сделать предложение</button>
          <button className="btn-cancel">Отменить и закрыть</button>
        </div>
      </div>

      {toast && <div className="toast">{toast}</div>}
    </div>
  );
};

export default TradePage;