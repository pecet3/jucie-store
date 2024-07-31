import pjlogo from "../assets/pjlogo.png"
import cart from "../assets/cart.png"
import { Link } from "react-router-dom"
export const Navbar = () => {

    return (
        <nav className="flex justify-between items-center p-4">
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
                    <a href="/juices" className="no-underline text-white px-2 font-semibold">JUICES</a>
                </li>
                <li className="list-none inline-block px-5">
                    <a href="/contact" className="no-underline text-white px-2 font-semibold">CONTACT</a>
                </li>
            </ul>
            <img src={cart} className="w-8 cursor-pointer" />
        </nav>
    )
}