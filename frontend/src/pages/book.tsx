import BookCover from "../assets/Book.png"

const Book = () => {
    return (
        <section>
            <div>
                <div>
                    <img src={BookCover} alt='book'/>
                </div>
                <div>
                    <h1 style={{ color: 'black' }}>Шантарам: Том 2</h1>
                    <h3 style={{ color: 'black'}}>Григори Дэвид Робертс</h3>
                </div>
            </div>
        </section>
    );
}

export default Book