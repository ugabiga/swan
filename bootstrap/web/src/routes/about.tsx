import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/about")({
    component: AboutComponent,
});

function AboutComponent() {
    return (
        <main className="w-screen h-screen">
            <h3 className="text-2xl ml-4 mt-4">
                About
            </h3>
        </main>
    )
}
