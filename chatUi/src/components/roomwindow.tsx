const RoomWindow = () => {
  return (
    <div className="w-full">
      <h1 className="text-center text-lg text-gray-800 px-4 py-2.5">
        RoomWIndow
      </h1>

      <div className="grid grid-cols-3 grid-rows-5 gap-4">
        <div className="col-span-2 row-span-2">
          <section className="mx-auto px-4 my-4 col-span-1">
            <h1 className="text-center text-xl font-semibold">Create Romm </h1>
            <form className="space-y-4 flex flex-col ">
              <div className="space-y-2">
                <input
                  type="text"
                  className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background"
                  placeholder="Text"
                />
              </div>

              <div className="space-y-2">
                <input
                  type="text"
                  className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background"
                  placeholder="Text"
                />
              </div>

              <div className="space-y-2">
                <button
                  type="submit"
                  className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm w-full sm:w-auto px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800 cursor-pointer"
                >
                  Submit
                </button>
              </div>
            </form>
          </section>
        </div>
        <div className="col-span-3 row-span-2 row-start-3">
          <section>
          </section>
        </div>
      </div>
    </div>
  );
};

export default RoomWindow;
