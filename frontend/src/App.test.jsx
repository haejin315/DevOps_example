import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import App from "./App";

// fetch를 mock 처리
beforeEach(() => {
  global.fetch = jest.fn(() =>
    Promise.resolve({
      ok: true,
      json: () => Promise.resolve({ A: 2, B: 3, Sum: 5 }),
    })
  );
});

afterEach(() => {
  global.fetch.mockClear();
});

test("덧셈 요청을 보내고 결과를 렌더링한다", async () => {
  render(<App />);

  fireEvent.change(screen.getByLabelText(/A:/i), { target: { value: "2" } });
  fireEvent.change(screen.getByLabelText(/B:/i), { target: { value: "3" } });

  fireEvent.click(screen.getByRole("button", { name: /Add/i }));

  await waitFor(() => {
    expect(screen.getByText("Result: 2 + 3 = 5")).toBeInTheDocument();
  });
});
