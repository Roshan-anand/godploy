"use client";

import { ModeToggle } from "@/components/mode-toggle";
import { Button } from "@hasu/ui/components/button";

export default function Home() {
  return (
    <div className="container mx-auto max-w-3xl px-4 py-2">
      <h1>Langind page</h1>
      <ModeToggle />
      <Button variant={"destructive"}>click</Button>
    </div>
  );
}
