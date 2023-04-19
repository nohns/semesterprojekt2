
int main()
{

  MotorDriver motor;

  Control control(&motor);

  x10Driver x10(&control);

  UartDriver uart(&control);

  Button button(&control);

  for (;;)
  {
    x10.ProcessInput();
    uart.ProcessInput();
    button.ProcessInput();
  }

  return 0;
}