import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/about")({
    component: AboutComponent,
});

function AboutComponent() {
    return (
        <main className="p-4 space-y-4">
            <h3 className="text-2xl">About</h3>
        </main>
    )
}
