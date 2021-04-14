import javax.swing.*;
import javax.swing.event.ChangeEvent;
import javax.swing.event.ChangeListener;
import java.awt.*;


public class PictureForm {
    private JPanel mainPanel;
    private JSpinner spinner1;
    private JTextField textField1;

    public PictureForm() {
        spinner1.addChangeListener(new ChangeListener() {
            @Override
            public void stateChanged(ChangeEvent e) {
                int  radius = (int) spinner1.getValue();

                double area=3.14*radius*radius;
                textField1.setText(String.format("%.2f", area));
            }
        });
        spinner1.setValue(20) ;
    }

    public static void main(String[] args) {
        JFrame frame = new JFrame("Окружность");
        frame.setContentPane(new PictureForm().mainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }
}
