import { Link } from "react-router-dom"
import { useStoreContext } from "../utils/storeContext"

export const NavBasket = () => {
    const { basket, finalPrice, changeItemQuantity } = useStoreContext()
    return (
        <div className="absolute top-16 right-4 flex flex-col rounded-lg backdrop-blur-md
         bg-white bg-opacity-10 p-2 z-50 max-h-96">
            <div className="overflow-y-scroll">
                {basket.map(i => (
                    <div key={i.product.id + i.price + i.strength}
                        className="flex items-center gap-1"
                    >
                        <img src={i.product.image_url} className="w-16 h-16" />
                        <p>{i.product.name}</p>
                        <p>{i.strength} MG</p>
                        <p>{i.capacity} ML</p>
                        <span>
                            <button onClick={() => changeItemQuantity(i, i.quantity + 1)}>+</button>
                            {i.quantity}
                            <button onClick={() => changeItemQuantity(i, i.quantity - 1)}>-</button>
                        </span>
                    </div>
                ))}
            </div>
            {finalPrice}
            <Link to="/basket" >
                Order
            </Link>
        </div>
    )
}