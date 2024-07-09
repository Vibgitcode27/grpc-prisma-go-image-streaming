-- CreateTable
CREATE TABLE "Laptop" (
    "id" TEXT NOT NULL,
    "brand" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "cpu" TEXT NOT NULL,
    "gpu" TEXT NOT NULL,
    "ram" TEXT NOT NULL,
    "keyboard" TEXT NOT NULL,
    "storages" TEXT NOT NULL,
    "screen" TEXT NOT NULL,
    "weightKg" DOUBLE PRECISION,
    "weightLbs" TEXT,
    "priceInr" DOUBLE PRECISION NOT NULL,
    "releaseYear" INTEGER NOT NULL,
    "updatedAt" TIMESTAMP(3) NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "Laptop_id_key" ON "Laptop"("id");
