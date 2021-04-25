import javax.swing.*;
import java.awt.*;
import java.awt.geom.AffineTransform;

public class Rectangle {
    private JPanel mainPanel;
    private CanvasPanel canvasPanel;
    private JSpinner a;
    private JSpinner b;
    private JSpinner alfa;

    public Rectangle() {
        a.addChangeListener(e -> canvasPanel.setA((Integer) a.getValue()));
        b.addChangeListener(e -> canvasPanel.setB((Integer) b.getValue()));
        alfa.addChangeListener(e -> canvasPanel.setAlfa((Integer) alfa.getValue()));

        a.setValue(500);
        b.setValue(300);
        alfa.setValue(5);
        canvasPanel.setA((Integer) a.getValue());
        canvasPanel.setB((Integer) b.getValue());
        canvasPanel.setAlfa((Integer) alfa.getValue());

    }
    public static void main(String[] args) {
        JFrame frame = new JFrame("Прямоугольник");
        frame.setContentPane(new Rectangle().mainPanel);
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.pack();
        frame.setVisible(true);
    }
    private void createUIComponents() {
        canvasPanel=new CanvasPanel();
    }
}
class CanvasPanel extends JPanel {
    private int a,b,alfa;
    public void setA(int a) {this.a=a; repaint();}
    public void setB(int b) {this.b=b; repaint();}
    public void setAlfa(int alfa) {this.alfa=alfa; repaint();}
    public void paint(Graphics g) {
        Graphics2D g2d = (Graphics2D) g;
        AffineTransform tx = new AffineTransform();
        tx.rotate(-1*alfa*Math.PI/180);
        java.awt.Rectangle shape = new java.awt.Rectangle(100, 100, a, b);
        Shape newShape = tx.createTransformedShape(shape);
g2d.clearRect(0,0,1000,1000);
        g2d.draw(newShape);
    }
}
