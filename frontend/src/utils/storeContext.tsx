import React, { createContext, useContext, useState, ReactNode } from 'react';
import { Price, Product } from './types';


type StoreContextType = {
    products: Product[]
    addProducts: (products: Product[]) => void
    getProductById: (id: number) => Product | undefined
    prices: Price[]
    addPrices: (prices: Price[]) => void
}

const StoreContext = createContext<StoreContextType | undefined>(undefined);

export const StoreProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
    const [products, setProducts] = useState<Product[]>([]);
    const addProducts = (products: Product[]) => {
        setProducts(products);
    };
    const getProductById = (id: number) => {
        const product = products.find(p => p.id === id)
        console.log(product)
        return product
    }

    const [prices, setPrices] = useState<Price[]>([])
    const addPrices = (prices: Price[]) => {
        setPrices(prices)
    }

    return (
        <StoreContext.Provider value={{ products, addProducts, getProductById, prices, addPrices }}>
            {children}
        </StoreContext.Provider>
    );
};

export const useStoreContext = () => {
    const context = useContext(StoreContext);
    if (context === undefined) {
        throw new Error('useStoreContext must be used within a StoreProvider');
    }
    return context;
};