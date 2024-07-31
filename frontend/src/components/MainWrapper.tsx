import React from "react"

export const MainWrapper = ({ children }: { children: React.ReactNode }) => {
    return (
        <main className="flex justify-center flex-col items-center">
            {children}
        </main>
    )
}