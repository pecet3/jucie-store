import pjlogo from "../assets/pjlogo.png"
import cart from "../assets/cart.png"
import { Link } from "react-router-dom"
import { FaShoppingBasket } from "react-icons/fa"
import { useStoreContext } from "../utils/storeContext"
import { useState } from "react"
import { NavBasket } from "./NavBasket"
export const Navbar = () => {
    const { basket, itemsCount } = useStoreContext()
    const [isBasket, setIsBasket] = useState(false)
    return (
        <nav className="relative flex justify-between items-center p-4">
            <Link to="/" >
                <img src={pjlogo} className="w-40 cursor-pointer" />
            </Link>
            <ul className="flex-1 text-center">
                <li className="list-none inline-block px-5">
                    <Link to="/how"
                        className="no-underline text-white px-2 font-semibold">HOW ?
                    </Link>
                </li>
                <li className="list-none inline-block px-5">
                    <Link to="/juices" className="no-underline text-white px-2 font-semibold">JUICES</Link>
                </li>
                <li className="list-none inline-block px-5">
                    <Link to="/contact" className="no-underline text-white px-2 font-semibold">CONTACT</Link>
                </li>
            </ul>
            <button onClick={() => setIsBasket(prev => !prev)} className="relative">
                <FaShoppingBasket size={36} />
                {itemsCount > 0 && !isBasket
                    ? <span className="absolute top-7 bg-purple-800 text-white rounded-full  px-1.5 text-sm">
                        {itemsCount}
                    </span>
                    : null}
            </button>
            {itemsCount > 0 && isBasket
                ? <NavBasket />
                : null}
        </nav>
    )
}